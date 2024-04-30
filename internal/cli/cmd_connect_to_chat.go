package cli

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/mistandok/chat-client/internal/common_error"
	"github.com/mistandok/chat-client/internal/service"
	"github.com/spf13/cobra"
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
					c.writer.Warning(err.Error())
				}
				c.logger.Err(err).Msg("неудачное подключение к чату")

				return
			}
			c.writer.Info("успешное подключение к чату")

			go func() {
				c.processIncomingMessageFromStream(stream)
			}()

			c.processOutgoingMessageToClient(cmd.Context(), chatID)
		},
	}
}

func (c *Chat) processIncomingMessageFromStream(stream service.StreamReader) {
	for {
		message, errRecv := stream.Recv()
		if errRecv == io.EOF {
			return
		}
		if errRecv != nil {
			log.Println("ошибка во время считывания сообщений из stream: ", errRecv)
			return
		}

		c.writer.Message(message.CreatedAt, message.FromUserName, message.Text)
	}
}

func (c *Chat) processOutgoingMessageToClient(ctx context.Context, chatID int64) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		msg, err := c.writer.ScanMessage()
		if err != nil {
			c.writer.Info("выход из чата")
			break
		}
		c.writer.CleanPreviousLine()

		if err = c.chatService.SendMessage(ctx, chatID, msg, time.Now()); err != nil {
			c.writer.Error("не удалось отправить сообщение, работаем над решением проблемы")
		}
	}

	err := scanner.Err()
	if err != nil {
		log.Println("failed to scan message: ", err)
	}
}
