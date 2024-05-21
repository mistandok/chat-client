package env

import (
	"errors"
	"os"

	"github.com/mistandok/chat-client/internal/config"
)

const (
	chatHostEnvName = "CHAT_HOST"
	chatPortEnvName = "CHAT_PORT"
)

// ChatCfgSearcher searcher for grpc config.
type ChatCfgSearcher struct{}

// NewChatCfgSearcher get instance for grpc config searcher.
func NewChatCfgSearcher() *ChatCfgSearcher {
	return &ChatCfgSearcher{}
}

// Get searcher for grpc config.
func (s *ChatCfgSearcher) Get() (*config.GRPCConfig, error) {
	host := os.Getenv(chatHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("chat host not found")
	}

	port := os.Getenv(chatPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("chat port not found")
	}

	return &config.GRPCConfig{
		Host: host,
		Port: port,
	}, nil
}
