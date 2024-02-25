package app

import (
	"context"
	"github.com/kom1ssar/grpc-auth/internal/closer"
	"github.com/kom1ssar/grpc-auth/internal/config"
	"github.com/kom1ssar/grpc-auth/internal/interceptor"
	descAuthV1 "github.com/kom1ssar/grpc-auth/pkg/auth_v1"
	descUserV1 "github.com/kom1ssar/grpc-auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load("local.env")
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	//tlsCredentials, err := credentials.NewServerTLSFromFile("service.pem", "service.pem")
	//if err != nil {
	//	log.Fatalf("Error parse TLS creds %s", err.Error())
	//}
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		//grpc.Creds(tlsCredentials),
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	)
	reflection.Register(a.grpcServer)

	descUserV1.RegisterUserV1Server(a.grpcServer, a.serviceProvider.UserImplementation(ctx))
	descAuthV1.RegisterAuthServiceServer(a.grpcServer, a.serviceProvider.AuthImplementation(ctx))

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server running on %s", a.serviceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}
	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil

}
