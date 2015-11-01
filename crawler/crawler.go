package crawler

import (
	"fmt"

	"github.com/jaytaylor/html2text"
	textapi "github.com/louiseGrandjonc/aylien_textapi_go"
	"github.com/picturapoesis/constants"
	"github.com/picturapoesis/managers/events"
	"github.com/picturapoesis/models"
	"github.com/picturapoesis/utils"

	"html"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func GetExhibitionLinkList(m models.Museum) ([]string, error) {
	results := []string{}
	var link string

	resp, err := http.Get(m.Place.AgendaURL)
	if err != nil {
		return results, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	bodyStr := string(body)

	exibitRegexValue := fmt.Sprintf("?P<value>(%s)", m.ExhibitionRegex)

	exibitExp := regexp.MustCompile(fmt.Sprintf("((%s))", exibitRegexValue))

	matches := exibitExp.FindAllStringSubmatch(bodyStr, -1)
	set := make(map[string]bool)
	for _, match := range matches {
		link = match[0]
		if set[link] {
			continue
		} else {
			foundURL := strings.TrimPrefix(match[0], "\"")
			foundURL = strings.TrimSuffix(foundURL, "\"")
			results = append(results, html.UnescapeString(foundURL))
			set[link] = true
		}
	}

	return results, nil
}

func CrawlEventURL(url string, lang string, baseURL string) (models.Event, error) {

	auth := textapi.Auth{constants.AYLIEN_ID, constants.AYLIEN_KEY}
	client, err := textapi.NewClient(auth, true)
	if err != nil {
		panic(err)
	}

	params := &textapi.ExtractParams{URL: utils.GetFullURL(url, baseURL), BestImage: true, Language: lang}
	article, err := client.Extract(params)

	if err != nil {
		return models.Event{}, err
	}

	eventObject := models.Event{
		URL:         url,
		Description: article.Article,
		Title:       article.Title,
		Image:       article.Image,
		Lang:        lang,
	}

	resp, err := http.Get(utils.GetFullURL(url, baseURL))

	if err != nil {
		return eventObject, err
	}
	defer resp.Body.Close()

	bodyStr, err := html2text.FromReader(resp.Body)
	if err != nil {
		fmt.Print(err)
		return eventObject, err
	}

	dates := utils.RetrieveDatesFromString(bodyStr, lang, article.Article[:50])

	if len(dates) > 0 {
		eventObject.DateBegin = dates[0]
		if len(dates) > 1 {
			eventObject.DateEnd = dates[1]
		}
	}

	return eventObject, nil
}

func CreateEventFromLinkList(m models.Museum, linkList []string) ([]models.Event, error) {

	existingURLS, err := events.FindExistingEventURLList(linkList)
	toCreateEvents := []models.Event{}
	toCreate := []string{}

	createEvents := false

	if err != nil {
		fmt.Print(err)
	}

	if len(existingURLS) != 0 {
		for _, url := range linkList {
			found := false
			for _, eUrl := range existingURLS {
				if url == eUrl {
					found = true
				}
			}

			if !found {
				createEvents = true
				toCreate = append(toCreate, url)
			}
		}
	} else {
		createEvents = true
		toCreate = linkList
	}

	if createEvents {
		for _, toCreateURL := range toCreate {
			event, err := CrawlEventURL(toCreateURL, m.Lang, m.Place.URL)
			if err == nil {
				event.Museum = m
				event.MuseumID = m.ID
				toCreateEvents = append(toCreateEvents, event)
			}
		}
		createdEvents, err := events.BulkCreate(toCreateEvents)

		if err != nil {
			return []models.Event{}, err
		}
		return createdEvents, nil
	}
	return []models.Event{}, nil
}
