package crawler

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"
	"youtubeCrawler/models"
	"youtubeCrawler/store"
)

const DefaultNumberOfGoRoutines = 5

//MyClient is a custom http client
var MyClient = &http.Client{
	Timeout: 10 * time.Second,
	Jar:     nil,
}

type Crawler struct {
	//numberOfIterations int
	//firstLink          string
	data         chan models.NextLink
	stopSignal   chan bool
	OutputTarget io.Closer
	wg           sync.WaitGroup
	nGoroutines  int
	StoreManager store.Manager
}

func New(storeManager *store.Manager) *Crawler {
	nRoutines := getNumberOfRoutines()
	return &Crawler{
		//	firstLink:          link,
		//	numberOfIterations: numOfIterations,
		data:        make(chan models.NextLink, 500),
		wg:          sync.WaitGroup{},
		nGoroutines: nRoutines,
		stopSignal:  make(chan bool, nRoutines),
		StoreManager: *storeManager,
	}
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

//getResponse does GET request to specified URI
func getResponse(httpMethod, baseUrl, urlSuffix string, customHttpClient *http.Client) *http.Response {
	uri := baseUrl + urlSuffix
	req, err := getHTTPRequest(httpMethod, uri)
	if err != nil {
		log.Printf("Failed to do HTTP Request, reason: %v", err)
	}
	res, err := customHttpClient.Do(req)
	if err != nil {
		log.Printf("Failed to get response from %v; reason: %v", uri, err)
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

func getNumberOfRoutines() int {
	n, b := os.LookupEnv("CRAWLER")
	if b {
		num, err := strconv.Atoi(n)
		if err != nil {
			fmt.Errorf("Failed to convert value %v of env 'CRAWLER' to int. Setting default value of %v", n, DefaultNumberOfGoRoutines)
			return DefaultNumberOfGoRoutines
		}
		return num
	}
	return DefaultNumberOfGoRoutines
}

// Crawl crawls through youTube
func (c *Crawler) crawl(id int) {
	var title string
	var err error
	var urlSuffix string
	for {
		select {
		case nextLink := <-c.data:
			fmt.Printf("Thread ID-%v Got Link from channel: [ID: %v], [title: %s], [link: %s], [number: %v]\n", id, nextLink.Id, nextLink.Title, nextLink.Link, nextLink.Number)
			if nextLink.Number > nextLink.NOfIterations {
				fmt.Printf("Stopped crawling for [ID: %v]; reached no. of iterations '%v' of '%v' on thread ID-%v ", nextLink.Id, nextLink.Number, nextLink.NOfIterations, id)
				break
			}
			c.StoreManager.StorePipe <- nextLink
			res := getResponse("GET", "http://www.youtube.com", nextLink.Link, MyClient)
			urlSuffix, title, err = parseNextVideoData(res)
			res.Body.Close()
			if err != nil {
				fmt.Println("Failed parseNextVideoData, reason: ", err)
			}
			c.data <- models.NextLink{Id: nextLink.Id, NOfIterations: nextLink.NOfIterations, Title: title, Link: urlSuffix, Number: nextLink.Number + 1}
		case <-c.stopSignal:
			c.wg.Done()
			return
		default:

		}
	}
}

func (c *Crawler) Run() {
	c.wg.Add(c.nGoroutines)

	for i := 0; i < c.nGoroutines; i++ {
		fmt.Printf("Starting routine no. %v\n", i)
		go c.crawl(i)
	}
	go c.StoreManager.StoreData()
	c.wg.Wait()
	close(c.data)
	close(c.StoreManager.StorePipe)
	fmt.Println("All channels closed")
}

func (c *Crawler) Stop() {
	for i := 0; i < c.nGoroutines; i++ {
		c.stopSignal <- true
	}
}

func (c *Crawler) Add(firstLink models.NextLink) {
	c.data <- firstLink
}
