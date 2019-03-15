package crawler

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"youtubeCrawler/config"
	"youtubeCrawler/models"
	"youtubeCrawler/parsers"
	"youtubeCrawler/store"
)

//myClient is a custom http client
var myClient = &http.Client{
	Timeout: 10 * time.Second,
	Jar:     nil,
	Transport: &http.Transport{
		MaxIdleConns:15,
		IdleConnTimeout: 30 * time.Second,
	},
}

// Crawler struct holds all data needed for crawling
type Crawler struct {
	data          chan models.NextLink //chan used for crawling
	stopSignal    chan bool            //chan to stop all crawling threads
	wg            sync.WaitGroup       //crawling threads waitGroup
	StoreManager  *store.Manager        // manager for data storing
	Configuration config.CrawlerConfig
	parser        parsers.DataParser
}

// returns *Crawler
func New(storeManager *store.Manager, config config.CrawlerConfig, parser parsers.DataParser) *Crawler {

	return &Crawler{
		data: make(chan models.NextLink, 500),
		wg:   sync.WaitGroup{},
		//nGoroutines:  config.NumOfGoroutines,
		stopSignal:    make(chan bool, config.NumOfGoroutines),
		StoreManager:  storeManager,
		Configuration: config,
		parser:        parser,
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

// Crawl crawls through youTube
// takes data from Crawler.Data chan in form of nextLink struct
// checks if enough iterations has been done
// sends copy to Crawler.StoreManager.StorePipe to store data
// calls getResponse to get *http.Body used to call parseNextVideoData to get urlSuffix and title
// makes new NextLink struct and sends it to Crawler.Data chan to keep crawling
// if receives stopSignal, crawling for that given thread stops
//TODO BUG: if chan buffer value is smaller than number of links to crawl, threads get stuck
func (c *Crawler) crawl(id int, parser parsers.DataParser) {
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
			res := getResponse("GET", nextLink.BaseUrl, nextLink.Link, myClient)
			urlSuffix, title, err = c.parser.ParseData(res)
			res.Body.Close()
			if err != nil {
				fmt.Println("Failed parseNextVideoData, reason: ", err)
			}
			c.data <- models.NextLink{Id: nextLink.Id, NOfIterations: nextLink.NOfIterations, Title: title, Link: urlSuffix, Number: nextLink.Number + 1, BaseUrl: nextLink.BaseUrl}
		case <-c.stopSignal:
			c.wg.Done()
			fmt.Printf("Thread ID-%v received stop signal and stopped\n", id)
			return
		default:

		}
	}
}

//Starts crawling
func (c *Crawler) Run() {
	c.wg.Add(c.Configuration.NumOfGoroutines)

	for i := 0; i < c.Configuration.NumOfGoroutines; i++ {
		fmt.Printf("Starting routine no. %v\n", i+1)
		go c.crawl(i, parsers.YoutubeParser{})
	}
	go c.StoreManager.StoreData()
	c.wg.Wait()
	close(c.data)
	fmt.Println("c.data Closed")
	close(c.StoreManager.StorePipe)
	fmt.Println("c.StoreManager.StorePipe closed")
	fmt.Println("All channels closed")
}

// stops all crawling threads
func (c *Crawler) Stop() {
	for i := 0; i < c.Configuration.NumOfGoroutines; i++ {
		fmt.Printf("Sending stop signal to thread ID-%v\n", i)
		c.stopSignal <- true
	}
}

//Adds link to the Crawler.Data chan to crawl
func (c *Crawler) Add(firstLink models.NextLink) {
	c.data <- firstLink
}
