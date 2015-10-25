package models

import (
	"github.com/picturapoesis/databases/fields"
	"time"
)

var (
	valueIndex int
)

type Place struct {
	ID          int64
	Name        string
	Infos       string
	URL         string
	AgendaURL   string
	LastWatched time.Time
}

type Museum struct {
	ID      int64
	PlaceID int64

	Place           Place
	Schedule        fields.MultiTimeSlice
	ExhibitionRegex string
	Lang            string
}
