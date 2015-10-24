package fields

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type MultiTimeSlice [][]time.Time

// Implements sql.Scanner for the String slice type
// Scanners take the database value (in this case as a byte slice)
// and sets the value of the type. Here we cast to a string and
// do a regexp based parse

// Need to find any array and special values in arrays

var (
	unquotedChar  = `[^",\\{}]`
	unquotedValue = fmt.Sprintf("(%s)+", unquotedChar)

	quotedChar  = `[^"\\]|\\"|\\\\`
	quotedValue = fmt.Sprintf("\"(%s)*\"", quotedChar)

	arrayValue     = fmt.Sprintf("(?P<value>(%s|%s))", unquotedValue, quotedValue)
	simpleArrayExp = regexp.MustCompile(fmt.Sprintf("((%s)(,)?)", arrayValue))

	timeChar = `[0-9]+:[0-9]+:[0-9]+`

	multiTimeArray      = fmt.Sprintf("(([{\"]+%s,%s[\"}]+)|([{\"]+NULL,NULL[\"}]+))", timeChar, timeChar)
	multiTimeArrayValue = fmt.Sprintf("(?P<value>)(%s)", multiTimeArray)
	multiTimeArrayExp   = regexp.MustCompile(fmt.Sprintf("((%s)(,)?)", multiTimeArray))

	quotedTimeValue      = fmt.Sprintf("\"(%s)\"", timeChar)
	simpleTimeArrayValue = fmt.Sprintf("(?P<value>(%s|%s))", unquotedValue, quotedTimeValue)
	simpleTimeArrayExp   = regexp.MustCompile(fmt.Sprintf("((%s))", simpleTimeArrayValue))

	valueIndex int
)

func parseSimpleTimeArray(array string) []time.Time {
	results := make([]time.Time, 0)

	matches := simpleTimeArrayExp.FindAllStringSubmatch(array, -1)

	for _, match := range matches {
		s := match[valueIndex]
		s = strings.Trim(s, "\"")
		t, err := time.Parse("15:04:05", s)
		if err != nil && s == "NULL" {
			break
		}
		results = append(results, t)
	}

	return results
}

func parseMultiTimeArray(array string) [][]time.Time {
	results := make([][]time.Time, 0)
	newArray := make([]time.Time, 0)
	if array[0] == '{' {
		array = array[1:]
	}
	if last := len(array) - 1; last >= 0 && array[last] == '}' {
		array = array[:last]
	}

	matches := multiTimeArrayExp.FindAllStringSubmatch(array, -1)

	for _, match := range matches {
		newArray = parseSimpleTimeArray(match[valueIndex])
		results = append(results, newArray)
	}
	return results
}

func (s *MultiTimeSlice) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []bytes"))
	}

	asString := string(asBytes)
	parsed := parseMultiTimeArray(asString)
	(*s) = MultiTimeSlice(parsed)
	return nil
}
