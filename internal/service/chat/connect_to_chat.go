package chat

import (
	"context"
	"github.com/mistandok/chat-client/internal/model"
	"github.com/mistandok/chat-client/internal/service"
	tokenUtils "github.com/mistandok/chat-client/internal/utils/token"
	"strconv"
)

func (s *Service) ConnectChat(ctx context.Context, chatID int64) (service.StreamReader, error) {
	tokens, userClaims, err := s.getTokensAndUserClaims(ctx)
	if err != nil {
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
