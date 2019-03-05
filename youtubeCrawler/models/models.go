package models

import (
	"fmt"
	"strings"
)

/* VideoData struct to hold data about video link
Title - title of the video
Link - part of the link: eg /watch?v=sWcXBRTGrWo
*/
type NextLink struct {
	Title         string `json:"title"`
	Link          string `json:"link"`
	Number        int    `json:"number"`
	Id            string `json:"id"`
	NOfIterations int    `json:"n_of_iterations"`
	Stop          bool
}

type Job struct {
	Id              int
	Name            string
	FirstVideoLink  string
	FirstVideoTitle string
	NumIterations   uint
	Progress        uint
	Finished        bool
}

func NewNextLink(firstLink string, numberOfIterations int) NextLink {
	fmt.Printf("Num of iterations is %v\n", numberOfIterations)
	if numberOfIterations == 0 {
		numberOfIterations = 100
	}
	title := strings.Split(firstLink, "=")
	return NextLink{
		Title:         "",
		Link:          firstLink,
		Number:        0,
		Id:            title[1],
		NOfIterations: numberOfIterations,
		Stop:          false,
	}
}
