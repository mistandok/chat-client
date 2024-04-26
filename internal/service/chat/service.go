package chat

import (
	"context"
	"errors"

	"github.com/mistandok/chat-client/internal/functional_error"
	"github.com/mistandok/chat-client/internal/repository"
	"github.com/rs/zerolog"

	"github.com/mistandok/chat-client/internal/client"
	"github.com/mistandok/chat-client/internal/model"
)

// Service ..
type Service struct {
	logger *zerolog.Logger

	userClient client.UserClient
	authClient client.AuthClient

	tokensRepo repository.TokensRepository
}

// NewService ..
func NewService(logger *zerolog.Logger, userClient client.UserClient, authClient client.AuthClient, tokensRepository repository.TokensRepository) *Service {
	return &Service{
		logger:     logger,
		userClient: userClient,
		authClient: authClient,
		tokensRepo: tokensRepository,
	}
}

// CreateUser ..
func (s *Service) CreateUser(ctx context.Context, userForCreate model.UserForCreate) error {
	if err := s.userClient.Create(ctx, userForCreate); err != nil {
		switch {
		case errors.Is(err, client.ErrUserAlreadyExists) || errors.Is(err, client.ErrTooLongPass):
			return functional_error.NewFunctionalError(err)
		default:
			return err
		}
	}

	return nil
}

// LoginUser ..
func (s *Service) LoginUser(ctx context.Context, email string, password string) error {
	tokens, err := s.authClient.Login(ctx, email, password)
	if err != nil {
		switch {
		case errors.Is(err, client.ErrUserNotFound) || errors.Is(err, client.ErrIncorrectAuthData):
			return functional_error.NewFunctionalError(err)
		default:
			return err
		}
	}

	return s.tokensRepo.Save(ctx, tokens)
}

// RefreshUserTokens ..
func (s *Service) RefreshUserTokens(_ context.Context, _ string) error {
	return nil
}
