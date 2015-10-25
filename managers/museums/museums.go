package museums

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/picturapoesis/constants"
	"github.com/picturapoesis/databases/fields"
	"github.com/picturapoesis/models"
	"time"
)

func GetPlace(i int) (models.Place, error) {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		constants.DB_USER, constants.DB_PASSWORD, constants.DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return models.Place{}, err
	}
	defer db.Close()

	var (
		id          int64
		name        string
		infos       string
		lastWatched time.Time
		url         string
		agendaURL   string
	)

	stmt, err := db.Prepare("SELECT id, name, infos, url, agenda_url, last_watched FROM place_place WHERE id=$1")

	if err != nil {
		return models.Place{}, err
	}

	err = stmt.QueryRow(i).Scan(
		&id,
		&name,
		&infos,
		&url,
		&agendaURL,
		&lastWatched)

	if err != nil {
		return models.Place{}, err
	}

	return models.Place{id, name, infos, url, agendaURL, lastWatched}, nil
}

func GetMuseum(i int) (models.Museum, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		constants.DB_USER, constants.DB_PASSWORD, constants.DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return models.Museum{}, err
	}
	defer db.Close()

	var (
		id              int64
		placeID         int64
		name            string
		infos           string
		lastWatched     time.Time
		url             string
		agendaURL       string
		schedule        fields.MultiTimeSlice
		exhibitionRegex string
		lang            string
	)

	stmt, err := db.Prepare("SELECT m.id, m.schedule, m.exhibition_regex, m.lang, p.id, p.name, p.url, p.agenda_url, p.last_watched, p.infos FROM place_museum m, place_place p WHERE m.place_id = p.id AND m.id=$1")

	if err != nil {
		return models.Museum{}, err
	}

	err = stmt.QueryRow(i).Scan(
		&id,
		&schedule,
		&exhibitionRegex,
		&lang,
		&placeID,
		&name,
		&url,
		&agendaURL,
		&lastWatched,
		&infos)

	if err != nil {
		return models.Museum{}, err
	}

	return models.Museum{
		ID:              id,
		PlaceID:         placeID,
		Schedule:        schedule,
		ExhibitionRegex: exhibitionRegex,
		Lang:            lang,
		Place: models.Place{
			ID:          placeID,
			Name:        name,
			LastWatched: lastWatched,
			URL:         url,
			AgendaURL:   agendaURL,
			Infos:       infos,
		}}, nil
}

func IsOpened(m models.Museum, day int) bool {
	if day < 0 || day > 6 {
		return false
	}
	if len(m.Schedule[day]) > 0 {
		return true
	}
	return false
}
