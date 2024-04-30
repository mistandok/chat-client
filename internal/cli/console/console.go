package console

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"os"
	"sync"
	"time"
)

type ConsoleWriter struct {
	mx      sync.Mutex
	scanner *bufio.Scanner
}

func NewConsoleWriter() *ConsoleWriter {
	return &ConsoleWriter{
		mx:      sync.Mutex{},
		scanner: bufio.NewScanner(os.Stdin),
	}
}

// Info ..
func (c *ConsoleWriter) Info(msg string) {
	newMsg := color.GreenString("[INFO]: ") + msg
	fmt.Println(newMsg)
}

// Warning ..
func (c *ConsoleWriter) Warning(msg string) {
	newMsg := color.YellowString("[WARN]: ") + msg
	fmt.Println(newMsg)
}

// Error ..
func (c *ConsoleWriter) Error(msg string) {
	newMsg := color.RedString("[ERR]: ") + msg
	fmt.Println(newMsg)
}

// Message ..
func (c *ConsoleWriter) Message(msgDateTime time.Time, userName string, msg string) {
	newMsg := color.YellowString(fmt.Sprintf("[%v]", msgDateTime.Format(time.DateTime)))
	newMsg += color.GreenString(fmt.Sprintf("[from: %s]: ", userName))
	newMsg += msg

	c.mx.Lock()
	fmt.Println(newMsg)
	c.mx.Unlock()
}

// ScanMessage ..
func (c *ConsoleWriter) ScanMessage() (string, error) {
	c.scanner.Scan()
	msg := c.scanner.Text()
	if msg == "exit" {
		return "", errors.New("выход из чата")
	}

	return msg, nil
}

// CleanPreviousLine ..
func (c *ConsoleWriter) CleanPreviousLine() {
	c.mx.Lock()
	c.mx.Unlock()
	fmt.Printf("\033[1A\033[K")
}
