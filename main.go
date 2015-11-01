package main

import (
	"github.com/picturapoesis/crawler"
	"github.com/picturapoesis/managers/museums"
)

func main() {

	ids := []int{1, 2}

	for _, i := range ids {
		m, _ := museums.GetMuseum(i)
		res, _ := crawler.GetExhibitionLinkList(m)
		crawler.CreateEventFromLinkList(m, res)
	}
}
