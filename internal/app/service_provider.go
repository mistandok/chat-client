package app

import (
	"context"
	"github.com/mistandok/chat-client/internal/repository"
	"github.com/mistandok/chat-client/internal/repository/token"
	"log"
	"os"

	"github.com/mistandok/chat-client/internal/client"
	"github.com/mistandok/chat-client/internal/client/auth"
	"github.com/mistandok/chat-client/internal/client/user"
	"github.com/mistandok/chat-client/internal/service"
	"github.com/mistandok/chat-client/internal/service/chat"
	"github.com/mistandok/chat-client/pkg/auth_v1"
	"github.com/mistandok/chat-client/pkg/chat_v1"
	"github.com/mistandok/chat-client/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/mistandok/chat-client/internal/config"
	"github.com/mistandok/chat-client/internal/config/env"
	"github.com/rs/zerolog"
)

type serviceProvider struct {
	authConfig *config.GRPCConfig
	chatConfig *config.GRPCConfig

	authGRPCClient auth_v1.AuthV1Client
	userGRPCClient user_v1.UserV1Client
	chatGRPCClient chat_v1.ChatV1Client

	authClient client.AuthClient
	userClient client.UserClient

	tokensRepo repository.TokensRepository

	chatService service.ChatService

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
	if s.authGRPCClient == nil {
		cfg := s.AuthConfig()
		conn, err := grpc.Dial(
			cfg.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("ошибка при установлении соединения с auth-сервисом: %v", err)
		}

		s.authGRPCClient = auth_v1.NewAuthV1Client(conn)
	}

	return s.authGRPCClient
}

// UserV1Client ..
func (s *serviceProvider) UserV1Client(_ context.Context) user_v1.UserV1Client {
	if s.userGRPCClient == nil {
		cfg := s.AuthConfig()
		conn, err := grpc.Dial(
			cfg.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("ошибка при установлении соединения с auth-сервисом: %v", err)
		}

		s.userGRPCClient = user_v1.NewUserV1Client(conn)
	}

	return s.userGRPCClient
}

// ChatV1Client ..
func (s *serviceProvider) ChatV1Client(_ context.Context) chat_v1.ChatV1Client {
	if s.chatGRPCClient == nil {
		cfg := s.ChatConfig()
		conn, err := grpc.Dial(
			cfg.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("ошибка при установлении соединения с chat-сервисом: %v", err)
		}

		s.chatGRPCClient = chat_v1.NewChatV1Client(conn)
	}

	return s.chatGRPCClient
}

func (s *serviceProvider) AuthClient(ctx context.Context) client.AuthClient {
	if s.authClient == nil {
		s.authClient = auth.NewClient(s.Logger(), s.AuthV1Client(ctx))
	}

	return s.authClient
}

func (s *serviceProvider) UserClient(ctx context.Context) client.UserClient {
	if s.userClient == nil {
		s.userClient = user.NewClient(s.Logger(), s.UserV1Client(ctx))
	}

	return s.userClient
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chat.NewService(s.Logger(), s.UserClient(ctx), s.AuthClient(ctx), s.TokensRepo(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) TokensRepo(_ context.Context) repository.TokensRepository {
	if s.tokensRepo == nil {
		s.tokensRepo = token.NewRepo("/tmp/user_tokens")
	}

	return s.tokensRepo
}

func setupZeroLog(logConfig *config.LogConfig) *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: logConfig.TimeFormat}
	logger := zerolog.New(output).With().Timestamp().Logger()
	logger = logger.Level(logConfig.LogLevel)
	zerolog.TimeFieldFormat = logConfig.TimeFormat

	return &logger
}
