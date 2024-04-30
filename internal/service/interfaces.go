package service

import (
	"context"
	"github.com/mistandok/chat-client/internal/model"
	"time"
)

// ChatService ..
type ChatService interface {
	CreateUser(ctx context.Context, userForCreate model.UserForCreate) error
	LoginUser(ctx context.Context, email string, password string) error
	RefreshUserTokens(ctx context.Context, refreshToken string) error
	ConnectChat(ctx context.Context, chatID int64) (StreamReader, error)
	SendMessage(ctx context.Context, chatID int64, text string, messageTime time.Time) error
	CreateChat(ctx context.Context) (int64, error)
}

// StreamReader ..
type StreamReader interface {
	Recv() (*model.Message, error)
	Context() context.Context
}
