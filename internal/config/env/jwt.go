package env

import (
	"errors"
	"github.com/kom1ssar/grpc-auth/internal/config"
	"os"
	"strconv"
)

const (
	jwtSecretEnvName         = "JWT_SECRET"
	jwtRefreshExpiredEnvName = "JWT_REFRESH_EXPIRED"
	jwtAccessExpiredEnvName
)

var _ config.JWTConfig = (*jwtConfig)(nil)

type jwtConfig struct {
	secret         string
	expiredRefresh int
	expiredAccess  int
}

func NewJWTConfig() (config.JWTConfig, error) {
	secret := os.Getenv(jwtSecretEnvName)
	if len(secret) == 0 {
		return nil, errors.New("jwt_secret env not found")
	}

	expiredRefreshString := os.Getenv(jwtRefreshExpiredEnvName)
	if len(expiredRefreshString) == 0 {
		return nil, errors.New("jwt_refresh_expired env not found")
	}
	expiredRefresh, err := strconv.Atoi(expiredRefreshString)
	if err != nil {
		return nil, errors.New("jwt_refresh_expired env parse to int err")
	}

	expiredAccessString := os.Getenv(jwtAccessExpiredEnvName)
	if len(expiredAccessString) == 0 {
		return nil, errors.New("jwt_access_expired env not found")
	}
	expiredAccess, err := strconv.Atoi(expiredRefreshString)
	if err != nil {
		return nil, errors.New("jwt_access_expired env parse to int err")
	}

	return &jwtConfig{
		secret:         secret,
		expiredRefresh: expiredRefresh,
		expiredAccess:  expiredAccess,
	}, nil

}

func (j *jwtConfig) Secret() string {
	return j.secret
}

func (j *jwtConfig) ExpiredRefresh() int {
	return j.expiredRefresh
}

func (j *jwtConfig) ExpiredAccess() int {
	return j.expiredAccess
}
