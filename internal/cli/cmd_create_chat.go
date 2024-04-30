package cli

import (
	"fmt"

	"github.com/mistandok/chat-client/internal/common_error"
	"github.com/spf13/cobra"
)

func (c *Chat) createCreateChatCmd() *cobra.Command {
	return &cobra.Command{
		Use:   chat,
		Short: createChatDesc,
		Run: func(cmd *cobra.Command, args []string) {
			chatID, err := c.chatService.CreateChat(cmd.Context())
			if err != nil {
				if common_error.IsCommonError(err) {
					c.writer.Warning(err.Error())
				}
				c.logger.Err(err).Msg(err.Error())

				return
			}

			c.writer.Info(fmt.Sprintf("Чат успешно создан. Его идентификатор: %d", chatID))
		},
	}
}
