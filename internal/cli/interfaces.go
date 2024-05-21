package cli

import "time"

// ExternalWriter ..
type ExternalWriter interface {
	Info(msg string)
	Warning(msg string)
	Error(msg string)
	Message(msgDateTime time.Time, userName string, msg string)
	ScanMessage() (string, error)
	CleanPreviousLine()
}
