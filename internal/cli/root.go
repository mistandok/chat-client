package cli

import (
	"context"
	"github.com/mistandok/chat-client/internal/service"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

const (
	appName        = "cli-chat"
	appDesc        = "cli утилита для чата"
	create         = "create"
	createDesc     = "позволяет создать пользователя или чат"
	user           = "user"
	createUserDesc = "создает нового пользователя"
	login          = "login"
	loginUserDesc  = "осуществляет log-in пользователя"
)

type Chat struct {
	chatService service.ChatService
	logger      *zerolog.Logger

	rootCmd       *cobra.Command
	createCmd     *cobra.Command
	createUserCmd *cobra.Command
	loginUserCmd  *cobra.Command
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
	c.loginUserCmd = c.createLoginUserCmd()
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

func (c *Chat) combineCommand() {
	c.rootCmd.AddCommand(c.createCmd)
	c.rootCmd.AddCommand(c.loginUserCmd)

	c.createCmd.AddCommand(c.createUserCmd)

	c.createUserCmd.Flags().StringP("username", "u", "", "имя пользователя")
	c.createUserCmd.Flags().StringP("email", "e", "", "email пользователя")
	c.createUserCmd.Flags().StringP("password", "p", "", "пароль пользователя")
	c.createUserCmd.MarkFlagsRequiredTogether("username", "email", "password")

	c.loginUserCmd.Flags().StringP("email", "e", "", "email пользователя")
	c.loginUserCmd.Flags().StringP("password", "p", "", "пароль пользователя")
	c.loginUserCmd.MarkFlagsRequiredTogether("email", "password")
}
