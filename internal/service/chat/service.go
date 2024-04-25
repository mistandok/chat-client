package chat

import (
	"context"
	"github.com/mistandok/chat-client/internal/repository"
	"github.com/rs/zerolog"

	"github.com/mistandok/chat-client/internal/client"
	"github.com/mistandok/chat-client/internal/model"
)

type Service struct {
	logger *zerolog.Logger

	userClient client.UserClient
	authClient client.AuthClient

	tokensRepo repository.TokensRepository
}

func NewService(logger *zerolog.Logger, userClient client.UserClient, authClient client.AuthClient, tokensRepository repository.TokensRepository) *Service {
	return &Service{
		logger:     logger,
		userClient: userClient,
		authClient: authClient,
		tokensRepo: tokensRepository,
	}
}

func (s *Service) CreateUser(ctx context.Context, userForCreate model.UserForCreate) error {
	return s.userClient.Create(ctx, userForCreate)
}

func (s *Service) LoginUser(ctx context.Context, email string, password string) error {
	tokens, err := s.authClient.Login(ctx, email, password)
	if err != nil {
		return err
	}

	return s.tokensRepo.Save(ctx, tokens)
}
func (s *Service) RefreshUserTokens(ctx context.Context, refreshToken string) error {
	return nil
}
