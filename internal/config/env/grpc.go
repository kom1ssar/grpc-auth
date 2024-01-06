package env

import (
	"errors"
	"github.com/kom1ssar/grpc-auth/internal/config"
	"net"
	"os"
)

var _ config.GRPCConfig = (*grpcConfig)(nil)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type grpcConfig struct {
	host string
	port string
}

func (g *grpcConfig) Address() string {
	return net.JoinHostPort(g.host, g.port)
}

func NewGRPCConfig() (config.GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc_host env not found")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc_port env not found")
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil

}
