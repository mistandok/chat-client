package model

import "time"

type Message struct {
	FromUserId   int64
	FromUserName string
	Text         string
	CreatedAt    time.Time
}
