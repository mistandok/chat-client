package auth

import (
	"context"

	"github.com/mistandok/chat-client/internal/client"
	"github.com/mistandok/chat-client/internal/model"
	"github.com/mistandok/chat-client/pkg/auth_v1"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	logger *zerolog.Logger
	client auth_v1.AuthV1Client
}

func NewClient(logger *zerolog.Logger, client auth_v1.AuthV1Client) *Client {
	return &Client{
		logger: logger,
		client: client,
	}
}

func (c *Client) Login(ctx context.Context, email string, password string) (*model.Tokens, error) {
	resp, err := c.client.Login(ctx, &auth_v1.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return nil, client.ErrUserNotFound
			case codes.InvalidArgument:
				return nil, client.ErrIncorrectAuthData
			}
		}
		c.logger.Err(err).Msg("auth rpc error")
		return nil, err
	}

	return &model.Tokens{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}, nil
}

func (c *Client) RefreshTokens(ctx context.Context, refreshToken string) (*model.Tokens, error) {
	resp, err := c.client.RefreshTokens(ctx, &auth_v1.RefreshTokensRequest{RefreshToken: refreshToken})
	if err != nil {
		return nil, err
	}

	return &model.Tokens{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}, nil
}
