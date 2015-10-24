package models

import (
	"github.com/picturapoesis/databases/fields"
	"time"
)

var (
	valueIndex int
)

type Museum struct {
	ID              int64
	Name            string
	LastWatched     time.Time
	Schedule        fields.MultiTimeSlice
	URL             string
	AgendaURL       string
	ExhibitionRegex string
}
