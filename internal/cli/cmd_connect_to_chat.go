package cli

import (
	"bufio"
	"fmt"
	"github.com/mistandok/chat-client/internal/cli/console"
	"github.com/mistandok/chat-client/internal/common_error"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"time"
)

func (c *Chat) createConnectToChatCmd() *cobra.Command {
	return &cobra.Command{
		Use:   connectToChat,
		Short: connectToChatDesc,
		Run: func(cmd *cobra.Command, args []string) {
			chatID, err := cmd.Flags().GetInt64("chat_id")
			if err != nil {
				c.logger.Fatal().Msg(fmt.Sprintf("не задан номер чата: %v", err))
			}

			stream, err := c.chatService.ConnectChat(cmd.Context(), chatID)
			if err != nil {
				if common_error.IsCommonError(err) {
					console.Warning(err.Error())
				}
				c.logger.Err(err).Msg("неудачное подключение к чату")

				return
			}

			go func() {
				for {
					message, errRecv := stream.Recv()
					if errRecv == io.EOF {
						return
					}
					if errRecv != nil {
						log.Println("failed to receive message from stream: ", errRecv)
						return
					}

					console.OutMessage(message.CreatedAt, message.FromUserName, message.Text)
				}
			}()

			scanner := bufio.NewScanner(os.Stdin)

			for {
				msg, err := console.ScanMessageAndCleanConsoleLine(scanner)
				if err != nil {
					console.Info("выход из чата")
					break
				}

				if err = c.chatService.SendMessage(cmd.Context(), chatID, msg, time.Now()); err != nil {
					console.Error("не удалось отправить сообщение, работаем над решением проблемы")
				}
			}

			err = scanner.Err()
			if err != nil {
				log.Println("failed to scan message: ", err)
			}
		},
	}
}
