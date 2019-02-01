package crawler

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"regexp"
	"time"
	"youtubeCrawler/models"
)

//MyClient is a custom http client
var MyClient = &http.Client{
	Timeout: 2 * time.Second,
	Jar:     nil,
}

//GetHTTPRequest returns *Request to do Do method with
func getHTTPRequest(method, uri string) (*http.Request, error) {
	httpMethod := method
	req, err := http.NewRequest(httpMethod, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/html; charset=utf-8")
	return req, nil
}

func getResponse(method, baseUrl, urlSuffix string, customHttpClient *http.Client) *http.Response {
	httpMethod := method
	uri := baseUrl + urlSuffix
	req, err := getHTTPRequest(httpMethod, uri)
	if err != nil {
		log.Printf("Failed to get HTTP Request, reason: %v", err)
	}
	res, err := customHttpClient.Do(req)
	if err != nil {
		log.Printf("Failed to get response from %v%v, reason: %v", baseUrl, urlSuffix, err)
	}

	return res
}

func parseNextVideoData(res *http.Response) (link, title string, err error) {
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
							//fmt.Printf("RegExp result: %v\n", matched)
							//fmt.Printf("%v\n", a)
							//fmt.Printf("link: %v\n", a.Val)
							link = a.Val
							needTitle = true
						}
					}
					if needTitle == true {
						if a.Key == "title" {
							//fmt.Printf("title: %v\n", a.Val)
							//fmt.Printf("----\n")
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

func Crawl(baseUrl, urlSuffix string, count int) chan structs.VideoData {
	var title string
	var err error

	var dataLine = make(chan structs.VideoData, 50)

	go func() {
		for i := 0; i < count; i++ {
			res := getResponse("GET", baseUrl, urlSuffix, MyClient)
			urlSuffix, title, err = parseNextVideoData(res)
			res.Body.Close()
			if err != nil {
				fmt.Println("Failed parseNextVideoData, reason: ", err)
			}
			dataLine <- structs.VideoData{Title: string(title),
				Link: urlSuffix}
			fmt.Printf("Found link %v with title %v\n", urlSuffix, title)
		}
	}()

	return dataLine
}
