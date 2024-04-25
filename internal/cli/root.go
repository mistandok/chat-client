package cli

import (
	"context"
	"errors"

	"github.com/mistandok/chat-client/internal/client"
	"github.com/mistandok/chat-client/internal/model"
	"github.com/mistandok/chat-client/internal/service"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

const (
	appName        = "cli-chat"
	appDesc        = "cli утилита для чата"
	create         = "create"
	createDesc     = "позволяет создать пользователя или чат"
	createUser     = "user"
	createUserDesc = "создает нового пользователя"
)

type Chat struct {
	chatService service.ChatService
	logger      *zerolog.Logger

	rootCmd       *cobra.Command
	createCmd     *cobra.Command
	createUserCmd *cobra.Command
}

func NewChat(logger *zerolog.Logger, chatService service.ChatService) *Chat {
	chat := &Chat{chatService: chatService, logger: logger}
	chat.initCommands()
	chat.combineCommand()

	return chat
}

func (c *Chat) Execute(ctx context.Context) error {
	err := c.rootCmd.ExecuteContext(ctx)
	if err != nil {
		c.logger.Err(err).Msg("ошибка во время выполнения команды")
		return err
	}

	return nil
}

func (c *Chat) initCommands() {
	c.rootCmd = c.createRootCmd()
	c.createCmd = c.createCreateCmd()
	c.createUserCmd = c.createCreateUserCmd()
}

func (c *Chat) createRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   appName,
		Short: appDesc,
	}
}

func (c *Chat) createCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   create,
		Short: createDesc,
	}
}

func (c *Chat) createCreateUserCmd() *cobra.Command {
	return &cobra.Command{
		Use:   createUser,
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

func (c *Chat) combineCommand() {
	c.rootCmd.AddCommand(c.createCmd)

	c.createCmd.AddCommand(c.createUserCmd)

	c.createUserCmd.Flags().StringP("username", "u", "", "имя пользователя")
	c.createUserCmd.Flags().StringP("email", "e", "", "email пользователя")
	c.createUserCmd.Flags().StringP("password", "p", "", "пароль пользователя")
	c.createUserCmd.MarkFlagsRequiredTogether("username", "email", "password")
}
