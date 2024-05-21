package chat

import (
	"context"

	"github.com/mistandok/chat-client/internal/client"
	"github.com/mistandok/chat-client/internal/model"
	"github.com/mistandok/chat-client/pkg/chat_v1"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Client ..
type Client struct {
	logger *zerolog.Logger
	client chat_v1.ChatV1Client
}

// NewClient ..
func NewClient(logger *zerolog.Logger, client chat_v1.ChatV1Client) *Client {
	return &Client{
		logger: logger,
		client: client,
	}
}

// ConnectChat ..
func (c *Client) ConnectChat(ctx context.Context, connectChatIn model.ConnectChatIn) (client.StreamReader, error) {
	stream, err := c.client.ConnectChat(ctx, &chat_v1.ConnectChatRequest{
		ChatId:   connectChatIn.ChatID,
		UserId:   connectChatIn.UserID,
		UserName: connectChatIn.UserName,
	})
	if err != nil {
		return nil, err
	}

	return NewStreamChatReader(stream), nil
}

// SendMessage ..
func (c *Client) SendMessage(ctx context.Context, chatID int64, message model.Message) error {
	if _, err := c.client.SendMessage(ctx, &chat_v1.SendMessageRequest{
		Message: &chat_v1.Message{
			FromUserId:   message.FromUserID,
			FromUserName: message.FromUserName,
			Text:         message.Text,
			CreatedAt:    timestamppb.New(message.CreatedAt),
		},
		ToChatId: chatID,
	}); err != nil {
		return err
	}

	return nil
}

// CreateChat ..
func (c *Client) CreateChat(ctx context.Context, userIDs []int64) (int64, error) {
	resp, err := c.client.Create(ctx, &chat_v1.CreateRequest{UserIDs: userIDs})
	if err != nil {
		return 0, err
	}

	return resp.Id, nil
}
