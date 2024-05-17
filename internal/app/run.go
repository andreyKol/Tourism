package app

import (
	"Tourism/internal/common/config"
	"Tourism/internal/common/swagger"
	"Tourism/internal/domain/ws"
	"Tourism/internal/handlers/http"
	"Tourism/internal/infrastructure/filesystem"
	"Tourism/internal/infrastructure/repository/postgresql"
	"Tourism/internal/usecase"
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
	"log/slog"
	stdhttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const filesDir = "files"

func (a *App) Run(ctx context.Context) error {
	db, err := pgxpool.New(ctx, config.DatabaseURL())
	if err != nil {
		return fmt.Errorf("creating pgxpool: %w", err)
	}
	if err = db.Ping(ctx); err != nil {
		return fmt.Errorf("pinging database: %w", err)
	}
	slog.Info("Connected to database")

	m, err := migrate.New("file://migrations", config.DatabaseURL())
	if err != nil {
		return fmt.Errorf("creating migration: %w", err)
	}
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("applying migrations: %w", err)
	}

	repo := postgresql.NewRepository(db)

	filesFS, err := filesystem.New(filesDir)
	if err != nil {
		return fmt.Errorf("creating images fs: %w", err)
	}
	authUseCase := usecase.NewAuthUseCase(repo)
	userUseCase := usecase.NewUserUseCase(repo, filesFS)
	wsUseCase := usecase.NewWsUseCase(repo)
	hub := ws.NewHub()

	handler := http.NewHandler(
		authUseCase,
		userUseCase,
		wsUseCase,
		hub,
	)
	go hub.Run()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use()
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3555", "http://127.0.0.1:3555", "http://localhost:3000", "http://127.0.0.1:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Session-Id"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)
	r.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)

	r.Post(`/api/v1/auth/sign-in`, handler.SignIn)
	r.Post(`/api/v1/auth/sign-out`, handler.SignUp)

	r.Get(`/api/v1/users/{userId}`, handler.GetUser)
	r.Patch(`/api/v1/users/update/{userId}`, handler.UpdateUser)
	r.Post(`/api/v1/users/image/{userId}`, handler.SetUserImage)

	r.Post("/api/v1/ws/createRoom", handler.CreateRoom)
	r.Get("/api/v1/ws/joinRoom/{roomId}", handler.JoinRoom)
	r.Get("/api/v1/ws/getRooms", handler.GetRooms)
	r.Get("/api/v1/ws/getClients/ByRoomID/{roomId}", handler.GetClientsByRoomID)
	r.Get("/api/v1/ws/getRooms/ByClientID", handler.GetRoomsByClientID)

	swagger.AddSwaggerRoutes(r)

	a.httpServer = &stdhttp.Server{
		Addr:    fmt.Sprintf(":%d", config.HttpPort()),
		Handler: http.HandlerFromMuxWithBaseURL(handler, r, "/api/v1"),
	}

	httpServerCh := make(chan error)
	go func() {
		httpServerCh <- a.httpServer.ListenAndServe()
	}()

	slog.Info(
		"Server is started",
		slog.String("addr", a.httpServer.Addr),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Info("Interrupt signal: " + s.String())
	case err = <-httpServerCh:
		slog.Error("Server stop signal: " + err.Error())
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer shutdownCancel()

	err = a.httpServer.Shutdown(shutdownCtx)
	if err != nil {
		slog.Error("Failed to shutdown the server: " + err.Error())
	}
	db.Close()
	slog.Info("Server has been shut down successfully")

	return nil
}
