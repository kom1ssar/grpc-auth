package app

import (
	"context"
	"github.com/kom1ssar/grpc-auth/internal/api/auth_v1"
	"github.com/kom1ssar/grpc-auth/internal/api/user_v1"
	"github.com/kom1ssar/grpc-auth/internal/client/db"
	"github.com/kom1ssar/grpc-auth/internal/client/db/pg"
	"github.com/kom1ssar/grpc-auth/internal/client/db/transaction"
	"github.com/kom1ssar/grpc-auth/internal/closer"
	"github.com/kom1ssar/grpc-auth/internal/config"
	"github.com/kom1ssar/grpc-auth/internal/config/env"
	"github.com/kom1ssar/grpc-auth/internal/managers"
	"github.com/kom1ssar/grpc-auth/internal/managers/password"
	"github.com/kom1ssar/grpc-auth/internal/repository"
	userRepo "github.com/kom1ssar/grpc-auth/internal/repository/user"
	"github.com/kom1ssar/grpc-auth/internal/service"
	authSrv "github.com/kom1ssar/grpc-auth/internal/service/auth"
	userSrv "github.com/kom1ssar/grpc-auth/internal/service/user"
	"log"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
	jwtConfig  config.JWTConfig

	dbClient           db.Client
	transactionManager db.TxManager
	userRepository     repository.UserRepository

	userService service.UserService
	authService service.AuthService

	passwordManager managers.PasswordManager

	userImplementation *user_v1.Implementation
	authImplementation *auth_v1.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PgConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPgConfig()
		if err != nil {
			log.Fatalf("eror pg_config init %+v", err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("eror grpc_config init %+v", err)
		}
		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) JWTConfig() config.JWTConfig {
	if s.jwtConfig == nil {
		cfg, err := env.NewJWTConfig()
		if err != nil {
			log.Fatalf("eror jwt_config init %+v", err)
		}
		s.jwtConfig = cfg
	}

	return s.jwtConfig
}

func (s *serviceProvider) DbClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PgConfig().DSN())
		if err != nil {
			log.Fatalf("eror pg_client init %+v", err)

		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("eror pg_client init ping %+v", err.Error())
		}

		closer.Add(func() error {
			cl.Close()
			return nil
		})
		s.dbClient = cl

	}

	return s.dbClient
}

func (s *serviceProvider) TransactionManager(ctx context.Context) db.TxManager {
	if s.transactionManager == nil {
		s.transactionManager = transaction.NewTransactionManager(s.DbClient(ctx).DB())
	}

	return s.transactionManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepo.NewUserRepository(s.DbClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userSrv.NewUserService(s.UserRepository(ctx), s.PasswordManager(ctx))
	}

	return s.userService
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authSrv.NewAuthService(s.UserRepository(ctx), s.PasswordManager(ctx), s.JWTConfig())
	}

	return s.authService
}

func (s *serviceProvider) UserImplementation(ctx context.Context) *user_v1.Implementation {
	if s.userImplementation == nil {
		s.userImplementation = user_v1.NewImplementation(s.UserService(ctx))
	}

	return s.userImplementation
}

func (s *serviceProvider) AuthImplementation(ctx context.Context) *auth_v1.Implementation {
	if s.authImplementation == nil {
		s.authImplementation = auth_v1.NewImplementation(s.AuthService(ctx))
	}

	return s.authImplementation
}

func (s *serviceProvider) PasswordManager(_ context.Context) managers.PasswordManager {
	if s.passwordManager == nil {
		s.passwordManager = password.NewPasswordManager()
	}

	return s.passwordManager
}
