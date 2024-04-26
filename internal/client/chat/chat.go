package chat

import (
	"context"
	"github.com/mistandok/chat-client/internal/client"
	"github.com/mistandok/chat-client/internal/model"
	"github.com/mistandok/chat-client/pkg/chat_v1"
	"github.com/rs/zerolog"
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

// StreamChatReader ..
type StreamChatReader struct {
	stream chat_v1.ChatV1_ConnectChatClient
}

func NewStreamChatReader(stream chat_v1.ChatV1_ConnectChatClient) *StreamChatReader {
	return &StreamChatReader{stream: stream}
}

// Recv ..
func (s *StreamChatReader) Recv() (*model.Message, error) {
	message, err := s.stream.Recv()
	if err != nil {
		return nil, err
	}

	return &model.Message{
		FromUserId:   message.FromUserId,
		FromUserName: message.FromUserName,
		Text:         message.Text,
		CreatedAt:    message.CreatedAt.AsTime(),
	}, nil
}

// Context ..
func (s *StreamChatReader) Context() context.Context {
	return s.stream.Context()
}
