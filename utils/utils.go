package utils

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	// "strings"
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

func RetrieveDatesFromString(s string, lang string, desc string) []time.Time {

	res := [][]string{}

	if lang == "fr" {
		res = HandleRegexDate(constants.DateRegexFR, s)
	} else {
		res = HandleRegexDate(constants.DateRegexFR, s)
	}

	if len(res) == 0 {
		return []time.Time{}
	}
	possibleDate := [][]string{}

	for _, date := range res {

		// dateStr := date[0]
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

		if valid {
			possibleDate = append(possibleDate, dateSlice)
			// fmt.Println(dateStr)
			// strAfterDate := strings.SplitAfterN(s, dateStr, 1)
			// fmt.Print(strings.Contains(strAfterDate[0], desc))
			// if strings.Contains(strAfterDate[0], desc) {
			// 	strAfterDesc := strings.SplitAfterN(strAfterDate[0], desc, 1)
			// 	distance := len(strAfterDate[0]) - len(strAfterDesc[0])
			// 	possibleDate[distance] = dateSlice
			// }
		}
	}

	if len(possibleDate) > 1 {
		// bestMatch := GetClosestDate(possibleDate)
		bestMatch := possibleDate[0]
		return GetDatesFromBestMatch(bestMatch, lang)
	}
	bestMatch := possibleDate[0]
	return GetDatesFromBestMatch(bestMatch, lang)
}

func GetDatesFromBestMatch(dateSlice []string, lang string) []time.Time {
	var day int
	var month int
	var year int

	day2 := -1
	month2 := -1
	year2 := -1

	switch len(dateSlice) {
	case 2:
		day, _ = strconv.Atoi(dateSlice[0])
		month = GetMonthValue(dateSlice[1], lang)
		year = time.Now().Year()
	case 3:
		day, _ = strconv.Atoi(dateSlice[0])
		month = GetMonthValue(dateSlice[1], lang)
		year, _ = strconv.Atoi(dateSlice[2])

	case 4:
		day, _ = strconv.Atoi(dateSlice[0])
		month = GetMonthValue(dateSlice[1], lang)
		year = time.Now().Year()

		day2, _ = strconv.Atoi(dateSlice[2])
		month2 = GetMonthValue(dateSlice[3], lang)
		year2 = year
	case 5:
		// year repeated only once
		if len(dateSlice[2]) == 4 {
			// year is this value
			year, _ = strconv.Atoi(dateSlice[2])
			day, _ = strconv.Atoi(dateSlice[0])
			month = GetMonthValue(dateSlice[1], lang)

			day2, _ = strconv.Atoi(dateSlice[3])
			month2 = GetMonthValue(dateSlice[4], lang)
			year2 = year

		} else {
			year, _ = strconv.Atoi(dateSlice[4])
			day, _ = strconv.Atoi(dateSlice[0])
			month = GetMonthValue(dateSlice[1], lang)

			day2, _ = strconv.Atoi(dateSlice[2])
			month2 = GetMonthValue(dateSlice[3], lang)
			year2 = year
		}
	case 6:
		year, _ = strconv.Atoi(dateSlice[2])
		day, _ = strconv.Atoi(dateSlice[0])
		month = GetMonthValue(dateSlice[1], lang)

		year2, _ = strconv.Atoi(dateSlice[5])
		day2, _ = strconv.Atoi(dateSlice[3])
		month2 = GetMonthValue(dateSlice[4], lang)
	}

	times := []time.Time{}

	times = append(times, time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))

	if day2 != -1 {
		times = append(times, time.Date(year2, time.Month(month2), day2, 0, 0, 0, 0, time.UTC))

	}

	return times
}

func GetMonthValue(month string, lang string) int {
	if len(month) <= 2 {
		value, _ := strconv.Atoi(month)
		return value
	}
	if lang == "fr" {
		for index, value := range constants.MonthFR {
			reg, _ := regexp.Compile(value)
			if reg.MatchString(month) {
				return index + 1
			}
		}
	}

	return -1
}

func GetClosestDate(dates map[int][]string) []string {
	minKey := -1

	for k, _ := range dates {
		if minKey == -1 {
			minKey = k
		} else if minKey > k {
			minKey = k
		}
	}

	return dates[minKey]
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
