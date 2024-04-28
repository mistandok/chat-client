package console

import (
	"fmt"
	"time"

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

// Error ..
func Error(msg string) {
	newMsg := color.RedString("[ERR]: ") + msg
	fmt.Println(newMsg)
}

func OutMessage(msgDateTime time.Time, userName string, msg string) {
	newMsg := color.YellowString(fmt.Sprintf("[%v]", msgDateTime.Format(time.DateTime)))
	newMsg += color.GreenString(fmt.Sprintf("[from: %s]: ", userName))
	newMsg += msg

	fmt.Println(newMsg)
}
