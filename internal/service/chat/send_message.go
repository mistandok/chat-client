package chat

import (
	"context"
	"strconv"
	"time"

	"github.com/mistandok/chat-client/internal/model"
	tokenUtils "github.com/mistandok/chat-client/internal/utils/token"
)

// SendMessage ..
func (s *Service) SendMessage(ctx context.Context, chatID int64, text string, messageTime time.Time) error {
	tokens, userClaims, err := s.getTokensAndUserClaims(ctx)
	if err != nil {
		return err
	}

	ctx = tokenUtils.OutgoingCtxWithAccessToken(ctx, tokens.AccessToken)

	userID, err := strconv.ParseInt(userClaims.UserID, 10, 64)
	if err != nil {
		s.logger.Err(err).Msg("ошибка при преобразовании userID к int64")
		return err
	}

	if err = s.chatClient.SendMessage(ctx, chatID, model.Message{
		FromUserID:   userID,
		FromUserName: userClaims.UserName,
		Text:         text,
		CreatedAt:    messageTime,
	}); err != nil {
		s.logger.Err(err).Msg("Ошибка во время отправки сообщения")
	}

	return nil
}
