package events

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
	"github.com/picturapoesis/constants"
	"github.com/picturapoesis/models"
)

func FindExistingEventURLList(urlList []string) ([]string, error) {

	var existingURLS []string

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		constants.DB_USER, constants.DB_PASSWORD, constants.DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return existingURLS, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT url FROM event_exhibition WHERE url in ($1)", strings.Join(urlList, ", "))
	if err != nil {
		return existingURLS, err
	}

	defer rows.Close()
	for rows.Next() {
		var eventURL string
		err = rows.Scan(&eventURL)
		if err == nil {
			existingURLS = append(existingURLS, eventURL)
		}
	}
	err = rows.Err()

	if err != nil {
		return existingURLS, err
	}

	return existingURLS, nil
}

func BulkCreate(events []models.Event) ([]models.Event, error) {

	results := []models.Event{}

	valueStrings := make([]string, 0, len(events))

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		constants.DB_USER, constants.DB_PASSWORD, constants.DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return results, err
	}
	defer db.Close()

	for _, event := range events {
		valueStr := fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s)",
			event.MuseumID,
			event.URL,
			event.DateBegin,
			event.DateEnd,
			event.Description,
			event.Title,
			event.Lang)
		valueStrings = append(valueStrings, valueStr)
	}

	test, err := db.Exec("INSERT INTO event_exhibition (museum_id, url, date_begin, date_end, description, title, lang) VALUES $1", strings.Join(valueStrings, ","))

	fmt.Print(test)

	if err != nil {
		return results, err
	}
	return events, nil
}
