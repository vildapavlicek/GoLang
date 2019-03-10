package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
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
	m := http.NewServeMux()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      m,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	storeManager := store.New(conf.StoreConfig)
	defer storeManager.StoreDestination.Close()

	monster := crawler.New(storeManager, conf.CrawlerConfig)
	go monster.Run()

	handlers.SetHandlers(m, monster, server)
	go startServer(server)

	//TODO graceful shutdown

	for {
		select {
		case <-storeManager.Shutdown:
			fmt.Println("Server shutting down")
			server.Shutdown(context.TODO())
			return
		default:
		}
	}

	os.Exit(1)
}

func startServer(s *http.Server) {
	fmt.Printf("Server listening at port: %v\n", s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Failed to start server. Reason: %v", err)
	}
}
