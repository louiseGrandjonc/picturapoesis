package events

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/picturapoesis/constants"
	"github.com/picturapoesis/models"
)

func FindExistingEventURLList(urlList []string) ([]string, error) {

	var existingURLS []string

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		constants.DB_USER, constants.DB_PASSWORD, constants.DB_NAME)
	db, err := sqlx.Open("postgres", dbinfo)

	if err != nil {
		return existingURLS, err
	}
	defer db.Close()

	query, args, err := sqlx.In("SELECT url FROM event_exhibition WHERE url IN (?);", urlList)

	query = db.Rebind(query)
	rows, err := db.Query(query, args...)

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

func BulkCreate(eventsToCreate []models.Event) ([]models.Event, error) {

	results := []models.Event{}

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		constants.DB_USER, constants.DB_PASSWORD, constants.DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return results, err
	}
	defer db.Close()

	txn, err := db.Begin()
	if err != nil {
		// log.Fatal(err)
		return results, err
	}

	stmt, err := txn.Prepare(pq.CopyIn(
		"event_exhibition",
		"museum_id",
		"url",
		"date_begin",
		"date_end",
		"description",
		"title",
		"lang"))

	if err != nil {
		return results, err
		// log.Fatal(err)
	}

	for _, ev := range eventsToCreate {
		_, err = stmt.Exec(
			ev.MuseumID,
			ev.URL,
			ev.DateBegin.Format("2006-01-02"),
			ev.DateEnd.Format("2006-01-02"),
			ev.Description,
			ev.Title,
			ev.Lang)
		if err != nil {
			return results, err
			// log.Fatal(err)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return results, err
		// log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		return results, err
		// log.Fatal(err)
	}

	err = txn.Commit()
	if err != nil {
		return results, err
		// log.Fatal(err)
	}

	return eventsToCreate, nil
}
