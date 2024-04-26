package token

import (
	"context"
	"encoding/json"
	"os"

	"github.com/mistandok/chat-client/internal/model"
	repoModel "github.com/mistandok/chat-client/internal/repository/token/model"
)

// Repo ..
type Repo struct {
	filePath string
}

// NewRepo ..
func NewRepo(filePath string) *Repo {
	return &Repo{filePath: filePath}
}

// Save ..
func (r *Repo) Save(_ context.Context, tokens *model.Tokens) (err error) {
	repoTokens := &repoModel.Tokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
	}()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(repoTokens)
	if err != nil {
		return err
	}

	return nil
}

// Get ..
func (r *Repo) Get(_ context.Context) (tokens *model.Tokens, err error) {
	repoTokens := &repoModel.Tokens{}

	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = file.Close()
	}()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(repoTokens)
	if err != nil {
		return nil, err
	}

	tokens = &model.Tokens{
		AccessToken:  repoTokens.AccessToken,
		RefreshToken: repoTokens.RefreshToken,
	}

	return tokens, nil
}
