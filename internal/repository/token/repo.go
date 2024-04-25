package token

import (
	"context"
	"encoding/json"
	"github.com/mistandok/chat-client/internal/model"
	repoModel "github.com/mistandok/chat-client/internal/repository/token/model"
	"os"
)

type Repo struct {
	filePath string
}

func NewRepo(filePath string) *Repo {
	return &Repo{filePath: filePath}
}

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
