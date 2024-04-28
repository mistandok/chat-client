package chat

import (
	"context"
	"github.com/mistandok/chat-client/internal/client"
	"github.com/mistandok/chat-client/internal/model"
	"github.com/mistandok/chat-client/pkg/chat_v1"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Client struct {
	logger *zerolog.Logger
	client chat_v1.ChatV1Client
}

func NewClient(logger *zerolog.Logger, client chat_v1.ChatV1Client) *Client {
	return &Client{
		logger: logger,
		client: client,
	}
}

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

func (c *Client) SendMessage(ctx context.Context, chatID int64, message model.Message) error {
	if _, err := c.client.SendMessage(ctx, &chat_v1.SendMessageRequest{
		Message: &chat_v1.Message{
			FromUserId:   message.FromUserId,
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
