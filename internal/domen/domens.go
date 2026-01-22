package domen

import "time"

type Message struct {
	Timestamp time.Time
	Body      string
}
