package museums

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/picturapoesis/constants"
	"github.com/picturapoesis/databases"
	"github.com/picturapoesis/models"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

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
		name            string
		lastWatched     time.Time
		url             string
		agendaURL       string
		schedule        fields.MultiTimeSlice
		exhibitionRegex string
	)

	stmt, err := db.Prepare("SELECT id, name, url, agenda_url, last_watch, schedule, exhibition_regex FROM place_museum, place_place WHERE place_museum.place_id = place_place.id AND place_museum.id=$1")

	if err != nil {
		return models.Museum{}, err
	}
	err = stmt.QueryRow(i).Scan(
		&id,
		&name,
		&url,
		&agendaURL,
		&lastWatched,
		&schedule,
		&exhibitionRegex)

	if err != nil {
		return models.Museum{}, err
	}

	return models.Museum{id, name, lastWatched, schedule, url, agendaURL, exhibitionRegex}, nil
}

func (m models.Museum) IsOpened(days ...int) bool {
	var day int
	if len(days) == 0 {
		now := time.Now()
		day = int(now.Weekday())
	} else if days[0] >= 0 && days[0] <= 6 {
		day = days[0]
	}

	if len(m.schedule[day]) > 0 {
		return true
	}
	return false
}

func (m models.Museum) GetExhibitionLinkList() []string {
	results := make([]string, 0)
	var link string

	resp, err := http.Get(m.AgendaURL)
	if err != nil {
		return results
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(body)
	bodyStr := string(body)

	exibitRegexValue := fmt.Sprintf("?P<value>(%s)", m.exhibition_regex)

	exibitExp := regexp.MustCompile(fmt.Sprintf("((%s))", exibitRegexValue))

	matches := exibitExp.FindAllStringSubmatch(bodyStr, -1)
	set := make(map[string]bool)

	for _, match := range matches {
		link = match[valueIndex]
		if set[link] {
			continue
		} else {
			results = append(results, match[valueIndex])
			set[link] = true
		}
	}

	return results
}
