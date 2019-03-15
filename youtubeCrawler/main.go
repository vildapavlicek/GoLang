package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"youtubeCrawler/config"
	"youtubeCrawler/crawler"
	"youtubeCrawler/handlers"
	"youtubeCrawler/parsers"
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

	stop := make(chan os.Signal, 1)
	go catchSignal(stop)

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

	monster := crawler.New(storeManager, conf.CrawlerConfig, parsers.YoutubeParser{}, os.Stdout)
	go monster.Run()

	handlers.SetHandlers(m, monster)
	go startServer(server)

	//TODO: BUG if threads get stuck, doesn't shutdown
	for {
		select {
		case <-storeManager.Shutdown:
			fmt.Println("Server shutting down")
			server.Shutdown(context.TODO())
			os.Exit(1)
		case <-stop:
			monster.Stop()
		default:
		}
	}
}

func startServer(s *http.Server) {
	fmt.Printf("Server listening at port: %v\n", s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Failed to start server. Reason: %v", err)
	}
}

func catchSignal(stopChan chan os.Signal) {
	signal.Notify(stopChan, os.Interrupt)
}
