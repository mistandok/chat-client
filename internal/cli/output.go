package cli

import (
	"fmt"

	"github.com/fatih/color"
)

// Info ..
func Info(msg string) {
	newMsg := color.GreenString("[INFO]: ") + msg
	fmt.Println(newMsg)
}

// Warning ..
func Warning(msg string) {
	newMsg := color.YellowString("[WARN]: ") + msg
	fmt.Println(newMsg)
}
