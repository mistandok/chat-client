package cli

import (
	"github.com/mistandok/chat-client/internal/cli/console"
	"github.com/mistandok/chat-client/internal/common_error"
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
				if common_error.IsCommonError(err) {
					console.Warning(err.Error())
				}
				c.logger.Err(err).Msg(err.Error())

				return
			}

			console.Info("авторизация успешно завершена")
		},
	}
}
