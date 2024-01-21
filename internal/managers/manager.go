package managers

import "context"

type PasswordManager interface {
	Hash(ctx context.Context, password string) (string, error)
	Compare(ctx context.Context, hashPassword string, password string) bool
}
