package cli

import (
	"errors"
	"github.com/mistandok/chat-client/internal/client"
	"github.com/spf13/cobra"
)

func (c *Chat) createLoginUserCmd() *cobra.Command {
	return &cobra.Command{
		Use:   login,
		Short: loginUserDesc,
		Run: func(cmd *cobra.Command, args []string) {
			email, err := cmd.Flags().GetString("email")
			if err != nil {
				c.logger.Fatal().Msg("не задан email пользователя")
			}
			password, err := cmd.Flags().GetString("password")
			if err != nil {
				c.logger.Fatal().Msg("не задан пароль")
			}

			err = c.chatService.LoginUser(cmd.Context(), email, password)
			if err != nil {
				if errors.Is(err, client.ErrUserNotFound) || errors.Is(err, client.ErrIncorrectAuthData) {
					c.logger.Warn().Msg(err.Error())
					return
				}
			}

			c.logger.Info().Msg("вы успешно залогинились")
		},
	}
}
