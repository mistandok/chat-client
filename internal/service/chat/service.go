package chat

import (
	"context"
	"errors"
	"github.com/mistandok/chat-client/internal/service"
	tokenUtils "github.com/mistandok/chat-client/internal/utils/token"
	"strconv"

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

// ConnectChat ..
func (s *Service) ConnectChat(ctx context.Context, chatID int64) (service.StreamReader, error) {
	tokens, err := s.tokensRepo.Get(ctx)
	if err != nil {
		s.logger.Err(err).Msg("ошибка при попытке получения токена из хранилища")
		return nil, err
	}

	userClaims, err := tokenUtils.GetUserClaims(tokens.AccessToken)
	if err != nil {
		s.logger.Err(err).Msg("ошибка при парсинге токена")
		return nil, err
	}

	ctx = tokenUtils.OutgoingCtxWithAccessToken(ctx, tokens.AccessToken)

	userID, err := strconv.ParseInt(userClaims.UserID, 10, 64)
	if err != nil {
		s.logger.Err(err).Msg("ошибка при преобразовании userID к int64")
		return nil, err
	}

	stream, err := s.chatClient.ConnectChat(ctx, model.ConnectChatIn{
		ChatID:   chatID,
		UserID:   userID,
		UserName: userClaims.UserName,
	})
	if err != nil {
		s.logger.Err(err).Msg("ошибка при соединении с чатом")
		return nil, err
	}

	return stream, err
}

// RefreshUserTokens ..
func (s *Service) RefreshUserTokens(_ context.Context, _ string) error {
	return nil
}
