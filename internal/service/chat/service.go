package chat

import (
	"context"

	"github.com/mistandok/chat-client/internal/client"
	"github.com/mistandok/chat-client/internal/model"
)

type Service struct {
	userClient client.UserClient
}

func NewService(userClient client.UserClient) *Service {
	return &Service{userClient: userClient}
}

func (s *Service) CreateUser(ctx context.Context, userForCreate model.UserForCreate) error {
	return s.userClient.Create(ctx, userForCreate)
}

func (s *Service) LoginUser(ctx context.Context, email string, password string) error {
	return nil
}
func (s *Service) RefreshUserTokens(ctx context.Context, refreshToken string) error {
	return nil
}
