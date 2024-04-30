package chat

import (
	"context"
	"errors"

	"github.com/mistandok/chat-client/internal/client"
	"github.com/mistandok/chat-client/internal/common_error"
)

// LoginUser ..
func (s *Service) LoginUser(ctx context.Context, email string, password string) error {
	tokens, err := s.authClient.Login(ctx, email, password)
	if err != nil {
		switch {
		case errors.Is(err, client.ErrUserNotFound) || errors.Is(err, client.ErrIncorrectAuthData):
			return common_error.NewCommonError(err.Error(), err)
		default:
			return err
		}
	}

	return s.tokensRepo.Save(ctx, tokens)
}
