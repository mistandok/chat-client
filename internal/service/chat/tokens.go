package chat

import (
	"context"
	"errors"

	"github.com/mistandok/chat-client/internal/common_error"
	"github.com/mistandok/chat-client/internal/model"
	"github.com/mistandok/chat-client/internal/repository"
	tokenUtils "github.com/mistandok/chat-client/internal/utils/token"
)

func (s *Service) getTokensAndUserClaims(ctx context.Context) (*model.Tokens, *model.UserClaims, error) {
	tokens, err := s.tokensRepo.Get(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrTokensNotFound) {
			return nil, nil, common_error.NewCommonError("Необходимо авторизоваться", err)
		}
		s.logger.Err(err).Msg("ошибка при попытке получения токена из хранилища")
		return nil, nil, err
	}

	userClaims, err := tokenUtils.GetUserClaims(tokens.AccessToken)
	if err != nil {
		s.logger.Err(err).Msg("ошибка при парсинге токена")
		return nil, nil, err
	}

	return tokens, userClaims, nil
}
