package usecase

import (
	"Tourism/internal/common/errors"
	"Tourism/internal/domain"
	"Tourism/internal/infrastructure/repository"
	"context"
	stderrors "errors"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"log/slog"
	"path/filepath"
)

//go:generate mockgen -source=user.go -destination=./mocks/user.go -package=mocks

type UserRepository interface {
	GetUser(ctx context.Context, id int64) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	SetUserImage(ctx context.Context, id int64, imageID string) error
}

type UserUseCase struct {
	filesDir string
	userRepo UserRepository
	fileRepo FileRepository
}

func NewUserUseCase(userRepo UserRepository, fileRepo FileRepository) *UserUseCase {
	return &UserUseCase{
		filesDir: "users",
		userRepo: userRepo,
		fileRepo: fileRepo,
	}
}

func (uc *UserUseCase) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	user, err := uc.userRepo.GetUser(ctx, id)
	if err != nil {
		if stderrors.Is(err, repository.ErrNotFound) {
			return nil, errors.NewNotFoundError("user was not found", "user")
		}
		return nil, fmt.Errorf("getting user %d: %w", id, err)
	}

	return user, nil
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, user *domain.User) error {
	err := uc.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("updating user %d: %w", user.ID, err)
	}

	return nil
}

func (uc *UserUseCase) SetUserImage(ctx context.Context, id int64, img []byte) error {
	if len(img) > 5<<20 {
		return fmt.Errorf("image is too big", "file")
	}

	mt := mimetype.Detect(img)
	if !mt.Is("image/jpeg") &&
		!mt.Is("image/png") &&
		!mt.Is("image/webp") {
		return fmt.Errorf("invalid image mime type", "file")
	}

	fileID := uuid.New().String()
	fileName := fileID + mt.Extension()

	user, err := uc.userRepo.GetUser(ctx, id)
	if err != nil {
		if stderrors.Is(err, repository.ErrNotFound) {
			return errors.NewNotFoundError("user was not found", "user")
		}
		return fmt.Errorf("getting user %d: %w", id, err)
	}

	err = uc.userRepo.SetUserImage(ctx, id, fileName)
	if err != nil {
		return fmt.Errorf("setting user %d image: %w", id, err)
	}

	err = uc.fileRepo.Save(ctx, img, filepath.Join(uc.filesDir, fileName))
	if err != nil {
		return fmt.Errorf("saving user image: %w", err)
	}

	if *user.ImageID != "" {
		err = uc.fileRepo.Delete(ctx, filepath.Join(uc.filesDir, *user.ImageID))
		if err != nil {
			slog.Error("Deleting old user image", slog.String("error", err.Error()))
		}
	}

	return nil
}
