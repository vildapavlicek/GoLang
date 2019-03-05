package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"youtubeCrawler/crawler"
	"youtubeCrawler/handlers"
	"youtubeCrawler/store"
)

var firstLink = "/watch?v=DT61L8hbbJ4"
var secondLink = "/watch?v=Q3oItpVa9fs"


func main() {

	storeManager := store.New()
	monster := crawler.New(storeManager)
	go monster.Run()
	defer storeManager.StoreDestination.Close()

	m := http.NewServeMux()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      m,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	handlers.SetHandlers(m, monster)
	fmt.Printf("Server listening at port: %v\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Failed to start server. Reason: %v", err)
	}
}
