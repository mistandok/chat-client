package service

import (
	"context"

	"github.com/mistandok/chat-client/internal/model"
)

type ChatService interface {
	CreateUser(ctx context.Context, userForCreate model.UserForCreate) error
	LoginUser(ctx context.Context, email string, password string) error
	RefreshUserTokens(ctx context.Context, refreshToken string) error
}
