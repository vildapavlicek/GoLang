package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
	"youtubeCrawler/config"
	"youtubeCrawler/crawler"
	"youtubeCrawler/handlers"
	"youtubeCrawler/store"
)

var firstLink = "/watch?v=DT61L8hbbJ4"
var secondLink = "/watch?v=Q3oItpVa9fs"

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load '.env' config file. All values will be set to default if not set as system environment variable")
	}
}

func main() {

	conf := config.New()
	storeManager := store.New(conf.StoreConfig)

	monster := crawler.New(storeManager, conf.CrawlerConfig)
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
	//TODO graceful shutdown
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Failed to start server. Reason: %v", err)
	}
}
