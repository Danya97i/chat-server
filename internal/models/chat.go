package models

import "time"

type Chat struct {
	Id        int64
	Title     string
	CreatedAt time.Time
	// UserEmails []string
}
