package chat

import (
	"context"
	tokenUtils "github.com/mistandok/chat-client/internal/utils/token"
	"strconv"
)

func (s *Service) CreateChat(ctx context.Context) (int64, error) {
	s.logger.Info().Msg("попытка создать чат")

	tokens, userClaims, err := s.getTokensAndUserClaims(ctx)
	if err != nil {
		return 0, err
	}

	ctx = tokenUtils.OutgoingCtxWithAccessToken(ctx, tokens.AccessToken)

	userID, err := strconv.ParseInt(userClaims.UserID, 10, 64)
	if err != nil {
		s.logger.Err(err).Msg("ошибка при преобразовании userID к int64")
		return 0, err
	}

	chatID, err := s.chatClient.CreateChat(ctx, []int64{userID})
	if err != nil {
		return 0, err
	}

	return chatID, nil
}
