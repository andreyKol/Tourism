package domain

import (
	"fmt"
	"regexp"
	"time"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"

	"Tourism/internal/common/errors"
)

var (
	phoneRegexp = regexp.MustCompile("^[+]?[(]?[0-9]{3}[)]?[-\\s.]?[0-9]{3}[-\\s.]?[0-9]{4,6}$")
	emailRegexp = regexp.MustCompile("[^@ \\t\\r\\n]+@[^@ \\t\\r\\n]+\\.[^@ \\t\\r\\n]+")
)

type User struct {
	ID                int64
	Name              string
	Phone             string
	PasswordEncrypted string
	CreatedAt         time.Time

	Surname    *string
	Patronymic *string
	Age        *int16
	Gender     *int16
	Email      *string
	ImageID    *string
	LastOnline *time.Time
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.NewInvalidInputError("name cannot be empty", "name")
	}
	if utf8.RuneCountInString(u.Name) > 30 {
		return errors.NewInvalidInputError("name too long", "name")
	}

	if !phoneRegexp.MatchString(u.Phone) {
		return errors.NewInvalidInputError("incorrect phone format", "phone")
	}

	return nil
}

func (u *User) EncryptPassword() error {
	data, err := bcrypt.GenerateFromPassword([]byte(u.PasswordEncrypted), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("encrypting password: %w", err)
	}

	u.PasswordEncrypted = string(data)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordEncrypted), []byte(password)) == nil
}
