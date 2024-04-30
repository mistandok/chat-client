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

type ChatClient interface {
	ConnectChat(ctx context.Context, connectChatIn model.ConnectChatIn) (StreamReader, error)
	SendMessage(ctx context.Context, chatID int64, message model.Message) error
	CreateChat(ctx context.Context, userIDs []int64) (int64, error)
}

// StreamReader ..
type StreamReader interface {
	Recv() (*model.Message, error)
	Context() context.Context
}
