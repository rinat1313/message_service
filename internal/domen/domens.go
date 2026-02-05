package domen

import "time"

type Message struct {
	ID        int64
	Timestamp time.Time
	Body      string
	IdChat    int64
}
