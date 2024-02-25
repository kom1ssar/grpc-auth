package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kom1ssar/grpc-auth/internal/config"
	"github.com/kom1ssar/grpc-auth/internal/managers"
	"github.com/kom1ssar/grpc-auth/internal/model"
	"github.com/kom1ssar/grpc-auth/internal/utils"
	"time"
)

var _ managers.JWTManager = (*manager)(nil)

type manager struct {
	config config.JWTConfig
}

func (m *manager) GenerateAccess(userId int64, email string, role int32) (string, error) {
	claims := model.UserJwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.config.SecondsExpiredAccess()).Unix(),
		},
		UserId: userId,
		Email:  email,
		Role:   role,
	}

	token, err := utils.GenerateToken(&claims, []byte(m.config.AccessSecret()))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m *manager) GenerateRefresh(userId int64) (string, error) {
	claims := model.UserJwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.config.SecondsExpiredRefresh()).Unix(),
		},
		UserId: userId,
	}

	token, err := utils.GenerateToken(&claims, []byte(m.config.RefreshSecret()))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m *manager) VerifyAccess(token string) (*model.UserJwtClaims, error) {
	return utils.VerifyToken(token, []byte(m.config.AccessSecret()))
}

func (m *manager) VerifyRefresh(token string) (*model.UserJwtClaims, error) {
	return utils.VerifyToken(token, []byte(m.config.RefreshSecret()))
}

func NewJWTManager(config config.JWTConfig) managers.JWTManager {
	return &manager{config: config}
}
