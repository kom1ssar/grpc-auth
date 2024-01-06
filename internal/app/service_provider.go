package app

import (
	"context"
	"github.com/kom1ssar/grpc-auth/internal/api/user"
	"github.com/kom1ssar/grpc-auth/internal/client/db"
	"github.com/kom1ssar/grpc-auth/internal/client/db/pg"
	"github.com/kom1ssar/grpc-auth/internal/closer"
	"github.com/kom1ssar/grpc-auth/internal/config"
	"github.com/kom1ssar/grpc-auth/internal/config/env"
	"github.com/kom1ssar/grpc-auth/internal/repository"
	userRepo "github.com/kom1ssar/grpc-auth/internal/repository/user"
	"github.com/kom1ssar/grpc-auth/internal/service"
	userSrv "github.com/kom1ssar/grpc-auth/internal/service/user"
	"log"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	userRepository repository.UserRepository

	userService service.UserService

	userImplementation *user.Implementation
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

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepo.NewUserRepository(s.DbClient(ctx))
	}
	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userSrv.NewUserService(s.UserRepository(ctx))
	}
	return s.userService
}

func (s *serviceProvider) UserImplementation(ctx context.Context) *user.Implementation {
	if s.userImplementation == nil {
		s.userImplementation = user.NewImplementation(s.UserService(ctx))
	}
	return s.userImplementation
}
