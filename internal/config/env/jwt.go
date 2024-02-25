package env

import (
	"errors"
	"github.com/kom1ssar/grpc-auth/internal/config"
	"os"
	"strconv"
	"time"
)

const (
	jwtAccessSecretEnvName   = "JWT_ACCESS_SECRET"
	jwtRefreshSecretEnvName  = "JWT_REFRESH_SECRET"
	jwtRefreshExpiredEnvName = "JWT_REFRESH_EXPIRED"
	jwtAccessExpiredEnvName
)

var _ config.JWTConfig = (*jwtConfig)(nil)

type jwtConfig struct {
	accessSecret          string
	refreshSecret         string
	secondsExpiredRefresh time.Duration
	secondsExpireAccess   time.Duration
}

func NewJWTConfig() (config.JWTConfig, error) {
	accessSecret := os.Getenv(jwtAccessSecretEnvName)
	if len(accessSecret) == 0 {
		return nil, errors.New("jwt_access_secret env not found")
	}

	refreshSecret := os.Getenv(jwtRefreshSecretEnvName)
	if len(refreshSecret) == 0 {
		return nil, errors.New("jwt_refresh_secret env not found")
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
		accessSecret:          accessSecret,
		refreshSecret:         refreshSecret,
		secondsExpiredRefresh: time.Second * time.Duration(expiredRefresh),
		secondsExpireAccess:   time.Second * time.Duration(expiredAccess),
	}, nil

}

func (j *jwtConfig) AccessSecret() string {
	return j.accessSecret
}

func (j *jwtConfig) RefreshSecret() string {
	return j.refreshSecret
}

func (j *jwtConfig) SecondsExpiredRefresh() time.Duration {
	return j.secondsExpiredRefresh
}

func (j *jwtConfig) SecondsExpiredAccess() time.Duration {
	return j.secondsExpireAccess
}
