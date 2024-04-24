package app

import (
	"context"
	"github.com/mistandok/chat-client/pkg/auth_v1"
	"github.com/mistandok/chat-client/pkg/chat_v1"
	"github.com/mistandok/chat-client/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"

	"github.com/mistandok/chat-client/internal/config"
	"github.com/mistandok/chat-client/internal/config/env"
	"github.com/rs/zerolog"
)

type serviceProvider struct {
	authConfig *config.GRPCConfig
	chatConfig *config.GRPCConfig

	authClient auth_v1.AuthV1Client
	userClient user_v1.UserV1Client
	chatClient chat_v1.ChatV1Client

	logger *zerolog.Logger
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// AuthConfig ..
func (s *serviceProvider) AuthConfig() *config.GRPCConfig {
	if s.authConfig == nil {
		searcher := env.NewAuthCfgSearcher()
		cfg, err := searcher.Get()
		if err != nil {
			log.Fatalf("не удалось получить auth config: %v", err)
		}

		s.authConfig = cfg
	}

	return s.authConfig
}

// ChatConfig ..
func (s *serviceProvider) ChatConfig() *config.GRPCConfig {
	if s.chatConfig == nil {
		searcher := env.NewChatCfgSearcher()
		cfg, err := searcher.Get()
		if err != nil {
			log.Fatalf("не удалось получить chat config: %v", err)
		}

		s.chatConfig = cfg
	}

	return s.chatConfig
}

// Logger ..
func (s *serviceProvider) Logger() *zerolog.Logger {
	if s.logger == nil {
		cfgSearcher := env.NewLogCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("не удалось получить pg config: %s", err.Error())
		}

		s.logger = setupZeroLog(cfg)
	}

	return s.logger
}

// AuthV1Client ..
func (s *serviceProvider) AuthV1Client(_ context.Context) auth_v1.AuthV1Client {
	if s.authClient == nil {
		cfg := s.AuthConfig()
		conn, err := grpc.Dial(
			cfg.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("ошибка при установлении соединения с auth-сервисом: %v", err)
		}

		s.authClient = auth_v1.NewAuthV1Client(conn)
	}

	return s.authClient
}

// UserV1Client ..
func (s *serviceProvider) UserV1Client(_ context.Context) user_v1.UserV1Client {
	if s.userClient == nil {
		cfg := s.AuthConfig()
		conn, err := grpc.Dial(
			cfg.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("ошибка при установлении соединения с auth-сервисом: %v", err)
		}

		s.userClient = user_v1.NewUserV1Client(conn)
	}

	return s.userClient
}

// ChatV1Client ..
func (s *serviceProvider) ChatV1Client(_ context.Context) chat_v1.ChatV1Client {
	if s.chatClient == nil {
		cfg := s.ChatConfig()
		conn, err := grpc.Dial(
			cfg.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("ошибка при установлении соединения с chat-сервисом: %v", err)
		}

		s.chatClient = chat_v1.NewChatV1Client(conn)
	}

	return s.chatClient
}

func setupZeroLog(logConfig *config.LogConfig) *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: logConfig.TimeFormat}
	logger := zerolog.New(output).With().Timestamp().Logger()
	logger = logger.Level(logConfig.LogLevel)
	zerolog.TimeFieldFormat = logConfig.TimeFormat

	return &logger
}
