package client

import (
	"context"

	"github.com/mistandok/chat-client/internal/model"
)

// AuthClient ..
type AuthClient interface {
	Login(ctx context.Context, email string, password string) (*model.Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (*model.Tokens, error)
}

// UserClient ..
type UserClient interface {
	Create(ctx context.Context, userForCreate model.UserForCreate) error
}
