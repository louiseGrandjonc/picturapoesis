package crawler

import (
	"fmt"

	textapi "github.com/AYLIEN/aylien_textapi_go"
	"github.com/jaytaylor/html2text"
	"github.com/picturapoesis/constants"
	"github.com/picturapoesis/managers/events"
	"github.com/picturapoesis/models"
	"github.com/picturapoesis/utils"

	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func GetExhibitionLinkList(m models.Museum) ([]string, error) {
	results := make([]string, 0)
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
			results = append(results, foundURL)
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

	params := &textapi.ExtractParams{URL: utils.GetFullURL(url, baseURL), BestImage: true}
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

	fmt.Println(utils.GetFullURL(url, baseURL))
	dates := utils.RetrieveDatesFromString(bodyStr, lang)

	fmt.Println("returned dates")
	fmt.Print(dates)

	if err != nil {
		fmt.Print(err)
		return eventObject, err
	}

	return eventObject, nil
}

func CreateEventFromLinkList(m models.Museum, linkList []string) ([]models.Event, error) {

	existingURLS, err := events.FindExistingEventURLList(linkList)

	toCreateEvents := []models.Event{}

	if err != nil {
		fmt.Print(err)
	}

	var toCreate []string

	if len(existingURLS) != 0 {
		for _, url := range linkList {
			found := false
			for _, eUrl := range existingURLS {
				if url == eUrl {
					found = true
				}
			}

			if !found {
				toCreate = append(toCreate, url)
			}
		}
	} else {
		toCreate = linkList
	}

	for _, toCreateURL := range toCreate {
		event, err := CrawlEventURL(toCreateURL, m.Lang, m.Place.URL)
		if err == nil {
			event.Museum = m
			event.MuseumID = m.ID
			toCreateEvents = append(toCreateEvents, event)
		}
	}

	// createdEvents, err := events.BulkCreate(toCreateEvents)

	// if err != nil {
	// 	return createdEvents, err
	// }

	return toCreateEvents, nil
}
