package env

import (
	"errors"
	"net"
	"os"
	"user-service/internal/config"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"
)

type httpConfig struct {
	host string
	port string
}

func NewHTTPConfig() (config.HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

func (c *httpConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
