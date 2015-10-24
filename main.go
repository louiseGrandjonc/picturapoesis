package main

import (
	"fmt"
	"github.com/picturapoesis/crawler"
	"github.com/picturapoesis/managers/museums"
)

func main() {
	i := 1
	m, _ := museums.GetMuseum(i)
	// now := time.Now()
	// day := int(now.Weekday())

	// fmt.Print(m.Place)
	// fmt.Print(m.Schedule)
	// fmt.Print(m.ExhibitionRegex)
	// fmt.Println(museums.IsOpened(m, day))

	// Test opened on mondays
	// fmt.Println(museums.IsOpened(m, 1))

	res, _ := crawler.GetExhibitionLinkList(m)
	fmt.Print(res)
	fmt.Print(m.Place.URL)
	events, _ := crawler.CreateEventFromLinkList(m, res)
	fmt.Print(events)
	fmt.Println("stop")
}
