package chat

import (
	"context"
	"errors"
	"github.com/mistandok/chat-client/internal/client"
	"github.com/mistandok/chat-client/internal/common_error"
	"github.com/mistandok/chat-client/internal/model"
)

// CreateUser ..
func (s *Service) CreateUser(ctx context.Context, userForCreate model.UserForCreate) error {
	if err := s.userClient.Create(ctx, userForCreate); err != nil {
		switch {
		case errors.Is(err, client.ErrUserAlreadyExists) || errors.Is(err, client.ErrTooLongPass):
			return common_error.NewCommonError(err.Error(), err)
		default:
			return err
		}
	}

	return nil
}
