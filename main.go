package main

import (
	"fmt"
	"github.com/picturapoesis/models"
)

func main() {
	i := 1
	museum := models.GetMuseum(i)

	fmt.Println(museum.IsOpened())

	// Test opened on mondays
	fmt.Println(museum.IsOpened(1))

	museum.GetExhibitionLinkList()
}
