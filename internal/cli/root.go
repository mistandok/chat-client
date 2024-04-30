package cli

import (
	"context"
	"github.com/mistandok/chat-client/internal/service"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

const (
	appName           = "cli-chat"
	appDesc           = "cli утилита для чата"
	create            = "create"
	createDesc        = "позволяет создать пользователя или чат"
	user              = "user"
	createUserDesc    = "создает нового пользователя"
	login             = "login"
	loginUserDesc     = "осуществляет log-in пользователя"
	connectToChat     = "connect-chat"
	connectToChatDesc = "присоединение к заданному чату"
	chat              = "chat"
	createChatDesc    = "создает новый чат"
)

// Chat ..
type Chat struct {
	chatService service.ChatService
	logger      *zerolog.Logger

	rootCmd          *cobra.Command
	createCmd        *cobra.Command
	createUserCmd    *cobra.Command
	loginUserCmd     *cobra.Command
	connectToChatCmd *cobra.Command
	createChatCmd    *cobra.Command

	writer ExternalWriter
}

// NewChat ..
func NewChat(logger *zerolog.Logger, chatService service.ChatService, writer ExternalWriter) *Chat {
	chat := &Chat{chatService: chatService, logger: logger, writer: writer}
	chat.initCommands()
	chat.combineCommand()

	return chat
}

// Execute ..
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
	c.connectToChatCmd = c.createConnectToChatCmd()
	c.createChatCmd = c.createCreateChatCmd()
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
	c.rootCmd.AddCommand(c.connectToChatCmd)

	c.createCmd.AddCommand(c.createUserCmd)
	c.createCmd.AddCommand(c.createChatCmd)

	c.createUserCmd.Flags().StringP("username", "u", "", "имя пользователя")
	c.createUserCmd.Flags().StringP("email", "e", "", "email пользователя")
	c.createUserCmd.Flags().StringP("password", "p", "", "пароль пользователя")
	c.createUserCmd.MarkFlagsRequiredTogether("username", "email", "password")

	c.loginUserCmd.Flags().StringP("email", "e", "", "email пользователя")
	c.loginUserCmd.Flags().StringP("password", "p", "", "пароль пользователя")
	c.loginUserCmd.MarkFlagsRequiredTogether("email", "password")

	c.connectToChatCmd.Flags().Int64P("chat_id", "c", 0, "ID чата")
	if err := c.connectToChatCmd.MarkFlagRequired("chat_id"); err != nil {
		c.logger.Fatal().Msg(err.Error())
	}
}
