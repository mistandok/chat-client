package chat

import (
	"context"
	"github.com/mistandok/chat-client/internal/model"
	"github.com/mistandok/chat-client/pkg/chat_v1"
)

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
