package parsers

import (
	"errors"
	"golang.org/x/net/html"
	"net/http"
	"regexp"
)

type YoutubeParser struct {
}

type DataParser interface {
	ParseData(response *http.Response) (link, title string, err error)
}

func (y YoutubeParser) ParseData(res *http.Response) (link, title string, err error) {
	newLink, newTitle, err := parseYoutubeData(res)
	return newLink, newTitle, err
}

func parseYoutubeData(res *http.Response) (link, title string, err error) {
	needTitle := false
	tokenizer := html.NewTokenizerFragment(res.Body, `<div>`)

	for {
		tempTag := tokenizer.Next()
		switch {
		case tempTag == html.ErrorToken:
			return "", "", errors.New("EOF")
		case tempTag == html.StartTagToken:
			tag := tokenizer.Token()

			isAnchor := tag.Data == "a"
			if isAnchor {
				//fmt.Printf("We found a link no. %d\n tag is: %v\nand tag.Attr is: %v\n", i, tag, tag.Attr)
				for _, a := range tag.Attr {
					if a.Key == "href" {
						if matched, _ := regexp.MatchString(`/watch\?v=\w+`, a.Val); matched {
							link = a.Val
							needTitle = true
						}
					}
					if needTitle == true {
						if a.Key == "title" {
							title = a.Val
							needTitle = false
							return link, title, nil
						}
					}
				}
			}
		}
	}
}
