package cli

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"strings"
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
				c.logger.Fatal().Msg(err.Error())

				return
			}

			go func() {
				for {
					c.logger.Info().Msg("try get message")
					message, errRecv := stream.Recv()
					if errRecv == io.EOF {
						return
					}
					if errRecv != nil {
						log.Println("failed to receive message from stream: ", errRecv)
						return
					}

					log.Printf("[%v] - [from: %s]: %s\n",
						color.YellowString(message.CreatedAt.Format(time.RFC3339)),
						color.BlueString(message.FromUserName),
						message.Text,
					)
					c.logger.Info().Msg(message.Text)
				}
			}()

			scanner := bufio.NewScanner(os.Stdin)
			var lines strings.Builder

			for {
				scanner.Scan()
				line := scanner.Text()
				if len(line) == 0 {
					break
				}

				lines.WriteString(line)
				lines.WriteString("\n")
			}

			err = scanner.Err()
			if err != nil {
				log.Println("failed to scan message: ", err)
			}
		},
	}
}
