package console

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fatih/color"
)

func ScanMessageAndCleanConsoleLine(scanner *bufio.Scanner) (string, error) {
	fmt.Print(color.BlueString("[me]: "))
	scanner.Scan()
	msg := scanner.Text()
	if msg == "exit" {
		return "", errors.New("выход из чата")
	}

	cleanPreviousConsoleLine()

	return msg, nil
}

func cleanPreviousConsoleLine() {
	fmt.Printf("\033[1A\033[K")
}
