package models

import "time"

type Chat struct {
	Id        int64     `db:"id"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
}
