package console

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
)

// Writer ..
type Writer struct {
	mx      sync.Mutex
	scanner *bufio.Scanner
}

// NewConsoleWriter ..
func NewConsoleWriter() *Writer {
	return &Writer{
		mx:      sync.Mutex{},
		scanner: bufio.NewScanner(os.Stdin),
	}
}

// Info ..
func (c *Writer) Info(msg string) {
	newMsg := color.GreenString("[INFO]: ") + msg
	fmt.Println(newMsg)
}

// Warning ..
func (c *Writer) Warning(msg string) {
	newMsg := color.YellowString("[WARN]: ") + msg
	fmt.Println(newMsg)
}

// Error ..
func (c *Writer) Error(msg string) {
	newMsg := color.RedString("[ERR]: ") + msg
	fmt.Println(newMsg)
}

// Message ..
func (c *Writer) Message(msgDateTime time.Time, userName string, msg string) {
	newMsg := color.YellowString(fmt.Sprintf("[%v]", msgDateTime.Format(time.DateTime)))
	newMsg += color.GreenString(fmt.Sprintf("[from: %s]: ", userName))
	newMsg += msg

	c.mx.Lock()
	fmt.Println(newMsg)
	c.mx.Unlock()
}

// ScanMessage ..
func (c *Writer) ScanMessage() (string, error) {
	c.scanner.Scan()
	msg := c.scanner.Text()
	if msg == "exit" {
		return "", errors.New("выход из чата")
	}

	return msg, nil
}

// CleanPreviousLine ..
func (c *Writer) CleanPreviousLine() {
	c.mx.Lock()
	defer c.mx.Unlock()
	fmt.Printf("\033[1A\033[K")
}
