package user

import (
	"context"

	"github.com/mistandok/chat-client/internal/client"
	"github.com/mistandok/chat-client/internal/model"
	"github.com/mistandok/chat-client/pkg/user_v1"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	logger *zerolog.Logger
	client user_v1.UserV1Client
}

func NewClient(logger *zerolog.Logger, client user_v1.UserV1Client) *Client {
	return &Client{
		logger: logger,
		client: client,
	}
}

func (c *Client) Create(ctx context.Context, userForCreate model.UserForCreate) error {
	_, err := c.client.Create(ctx, &user_v1.CreateRequest{
		Name:            userForCreate.Name,
		Email:           userForCreate.Email,
		Password:        userForCreate.Password,
		PasswordConfirm: userForCreate.Password,
		Role:            user_v1.Role_USER,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.AlreadyExists:
				return client.ErrUserAlreadyExists
			case codes.InvalidArgument:
				return client.ErrIncorrectAuthData
			}
		}
		c.logger.Err(err).Msg("user rpc error")
		return err
	}

	return nil
}
