package cli

import (
	"errors"
	"github.com/mistandok/chat-client/internal/client"
	"github.com/mistandok/chat-client/internal/model"
	"github.com/spf13/cobra"
)

func (c *Chat) createCreateUserCmd() *cobra.Command {
	return &cobra.Command{
		Use:   user,
		Short: createUserDesc,
		Run: func(cmd *cobra.Command, args []string) {
			name, err := cmd.Flags().GetString("username")
			if err != nil {
				c.logger.Fatal().Msg("не задано имя пользователя")
			}
			email, err := cmd.Flags().GetString("email")
			if err != nil {
				c.logger.Fatal().Msg("не задан email пользователя")
			}
			password, err := cmd.Flags().GetString("password")
			if err != nil {
				c.logger.Fatal().Msg("не задан пароль")
			}

			err = c.chatService.CreateUser(cmd.Context(), model.UserForCreate{
				Name:     name,
				Email:    email,
				Password: password,
			})
			if err != nil {
				if errors.Is(err, client.ErrUserAlreadyExists) {
					c.logger.Warn().Msg(err.Error())
					return
				}
			}

			c.logger.Info().Msg("пользователь успешно создан")
		},
	}
}
