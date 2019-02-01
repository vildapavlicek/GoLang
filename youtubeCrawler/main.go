package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"youtubeCrawler/crawler"
	"youtubeCrawler/handlers"
	"youtubeCrawler/shared"
)

///watch?v=pVHKp6ffURY
var yt = "http://www.youtube.com"
var firstLink = "/watch?v=sWcXBRTGrWo"

func main() {

	db := shared.MyDB{
		User:   "root",
		Pwd:    "1111",
		DbUrl:  "tcp(127.0.0.1:3306)",
		DbName: "testdb",
	}

	db.OpenConnection()

	c := crawler.Crawl(yt, firstLink, 1000)

	db.TestInsertSuffixUrl(c)

	m := http.NewServeMux()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      m,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	handlers.SetHandlers(m)
	fmt.Printf("Server listening at port: %v", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Failed to start server. Reason: %v", err)
	}

}
