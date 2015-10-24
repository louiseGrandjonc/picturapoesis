package events

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
	"github.com/picturapoesis/constants"
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
	return []models.Event{}, nil
}
