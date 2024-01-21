package password

import (
	"context"
	desc "github.com/kom1ssar/grpc-auth/internal/managers"
	"golang.org/x/crypto/bcrypt"
)

var _ desc.PasswordManager = (*passwordManager)(nil)

type passwordManager struct {
}

func NewPasswordManager() desc.PasswordManager {
	return &passwordManager{}
}

func (p passwordManager) Hash(_ context.Context, password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (p passwordManager) Compare(_ context.Context, hashPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
