package repository

import (
	"context"

	"github.com/mistandok/chat-client/internal/model"
)

// TokensRepository ..
type TokensRepository interface {
	Save(ctx context.Context, tokens *model.Tokens) error
	Get(ctx context.Context) (*model.Tokens, error)
}
