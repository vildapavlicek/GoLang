package models

import (
	"fmt"
	"strings"
)

//VideoData struct to hold data about video link
type NextLink struct {
	Title         string `json:"title"` // Title of the video
	BaseUrl       string `json:"baseUrl"`
	Link          string `json:"link"`            // Link URL suffix `/watch?v=P-Xz-IeijSw`
	Number        int    `json:"number"`          // Number of iteration that data were received
	Id            string `json:"id"`              // ID string taken from Link in format `P-Xz-IeijSw`
	NOfIterations int    `json:"n_of_iterations"` // Number of link to crawl from origin (first link)
	Stop          bool   // Deceprated
}

// used to create first link to start crawling from
func NewNextLink(firstLink string, numberOfIterations int) NextLink {
	fmt.Printf("Num of iterations is %v\n", numberOfIterations)
	if numberOfIterations == 0 {
		numberOfIterations = 100
	}
	title := strings.Split(firstLink, "=")
	return NextLink{
		Title:         "",
		Link:          firstLink,
		BaseUrl:       "http://www.youtube.com",
		Number:        0,
		Id:            title[1],
		NOfIterations: numberOfIterations,
		Stop:          false,
	}
}
