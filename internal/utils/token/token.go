package token

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/mistandok/chat-client/internal/model"
	"google.golang.org/grpc/metadata"
)

func OutgoingCtxWithAccessToken(ctx context.Context, accessToken string) context.Context {
	outgoingMd := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	return metadata.NewOutgoingContext(ctx, outgoingMd)
}

func GetUserClaims(tokenString string) (*model.UserClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &model.UserClaims{})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*model.UserClaims)
	if !ok {
		return nil, errors.New("некорректное тело токена")
	}

	return claims, nil
}
