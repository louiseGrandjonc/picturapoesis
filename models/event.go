package models

import (
	"time"
)

type Event struct {
	ID       int64
	MuseumID int64

	Museum      Museum
	URL         string
	DateBegin   time.Time
	DateEnd     time.Time
	Description string
	Title       string
	Lang        string
	Image       string
}
