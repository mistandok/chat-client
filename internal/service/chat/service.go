package chat

import (
	"context"
	"github.com/mistandok/chat-client/internal/repository"
	"github.com/rs/zerolog"

	"github.com/mistandok/chat-client/internal/client"
)

// Service ..
type Service struct {
	logger *zerolog.Logger

	userClient client.UserClient
	authClient client.AuthClient
	chatClient client.ChatClient

	tokensRepo repository.TokensRepository
}

// NewService ..
func NewService(
	logger *zerolog.Logger,
	userClient client.UserClient,
	authClient client.AuthClient,
	chatClient client.ChatClient,
	tokensRepository repository.TokensRepository,
) *Service {
	return &Service{
		logger:     logger,
		userClient: userClient,
		authClient: authClient,
		chatClient: chatClient,
		tokensRepo: tokensRepository,
	}
}

// RefreshUserTokens ..
func (s *Service) RefreshUserTokens(_ context.Context, _ string) error {
	return nil
}
