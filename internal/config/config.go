package config

import (
	"net"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

// LogConfigSearcher interface for serach Log config.
type LogConfigSearcher interface {
	Get() (*LogConfig, error)
}

// AuthConfigSearcher interface for search grpc config
type AuthConfigSearcher interface {
	Get() (*GRPCConfig, error)
}

// ChatConfigSearcher interface for search grpc config
type ChatConfigSearcher interface {
	Get() (*GRPCConfig, error)
}

// Load dotenv from path to env
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

// GRPCConfig grpc config.
type GRPCConfig struct {
	Host string
	Port string
}

// Address get address for grpc server.
func (cfg *GRPCConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

// LogConfig config for zerolog.
type LogConfig struct {
	LogLevel   zerolog.Level
	TimeFormat string
}
