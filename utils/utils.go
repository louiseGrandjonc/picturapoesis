package utils

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"github.com/picturapoesis/constants"
)

func GetFullURL(initialURL string, baseURL string) string {
	parsedURL, err := url.Parse(initialURL)

	if err != nil {
		return initialURL
	}

	if parsedURL.IsAbs() {
		return initialURL
	}

	return fmt.Sprint(baseURL, initialURL)
}

func RetrieveDatesFromString(s string, lang string) []time.Time {

	res := [][]string{}

	if lang == "fr" {
		res = HandleRegexDate(constants.DateRegexFR, s)
	} else {
		res = HandleRegexDate(constants.DateRegexFR, s)
	}

	for _, date := range res {

		dateStr := date[0]
		dateSlice := CleanDateSlice(date[1:])
		dateLength := len(dateSlice)

		valid := true

		switch dateLength {
		case 2:
			day := dateSlice[0]
			month := dateSlice[1]
			year := time.Now().Year()
			valid = DateIsValid(day, month, year)
		case 3:
			day := dateSlice[0]
			month := dateSlice[1]
			year, errConv := strconv.Atoi(dateSlice[2])

			if errConv != nil {
				valid = false
			} else {
				valid = DateIsValid(day, month, year)
			}
		case 4:
			day := dateSlice[0]
			month := dateSlice[1]
			year := time.Now().Year()
			valid = DateIsValid(day, month, year)

			day2 := dateSlice[0]
			month2 := dateSlice[1]
			valid = DateIsValid(day2, month2, year)
		case 5:
			// year repeated only once
			if len(dateSlice[2]) == 4 {
				// year is this value
				year, errConv := strconv.Atoi(dateSlice[2])
				if errConv != nil {
					valid = false
					break
				}
				valid = DateIsValid(dateSlice[0], dateSlice[1], year)
				valid = DateIsValid(dateSlice[3], dateSlice[4], year)
			} else {
				year, errConv := strconv.Atoi(dateSlice[4])
				if errConv != nil {
					valid = false
					break
				}
				valid = DateIsValid(dateSlice[0], dateSlice[1], year)
				valid = DateIsValid(dateSlice[2], dateSlice[3], year)
			}
		case 6:
			year, errConv := strconv.Atoi(dateSlice[2])
			if errConv != nil {
				valid = false
				break
			}
			valid = DateIsValid(dateSlice[0], dateSlice[1], year)

			year, errConv = strconv.Atoi(dateSlice[5])
			if errConv != nil {
				valid = false
				break
			}
			valid = DateIsValid(dateSlice[3], dateSlice[4], year)
		}
		fmt.Print(valid)
		fmt.Print(dateStr)
	}

	return []time.Time{}
}

func HandleRegexDate(regex string, s string) [][]string {
	reg, err := regexp.Compile(regex)
	if err != nil {
		fmt.Print(err)
		return [][]string{}
	}
	return reg.FindAllStringSubmatch(s, -1)
}

func DateIsValid(day string, month string, year int) bool {
	_, errConv := strconv.Atoi(day)
	if errConv != nil {
		return false
	}
	if year < 2000 {
		return false
	}

	if len(month) <= 2 {
		// it a digit
		monthInt, errConv := strconv.Atoi(month)
		if errConv != nil {
			return false
		}
		if monthInt > 12 || monthInt < 1 {
			return false
		}
	} else {
		reg, _ := regexp.Compile(constants.MonthRegex)
		return reg.MatchString(month)
	}
	return true
}

func CleanDateSlice(date []string) []string {
	res := []string{}
	for _, value := range date {
		if len(value) > 0 {
			res = append(res, value)
		}
	}
	return res
}
