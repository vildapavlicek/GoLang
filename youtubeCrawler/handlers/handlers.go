package handlers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"youtubeCrawler/crawler"
	"youtubeCrawler/models"
)

// registers all handlers with ServeMux
func SetHandlers(m *http.ServeMux, c *crawler.Crawler) {
	m.HandleFunc("/", index)
	m.HandleFunc("/api/v1/link", linkHandler(c))
	m.HandleFunc("/api/v1/stop", stopAll(c))
}

//TODO should be used for landing page, so far used for testing tamplates
func index(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./views/index.gohtml"))
	tpl.Execute(w, "Vilda")
}

// accepts POST method to add new link for crawling if successful returns StatusCreated - 201 else StatusBadRequest 400
// GET method returns http.StatusMethodNotAllowed - 405
// default response set to http.StatusInternalServerError - 500
func linkHandler(crawler *crawler.Crawler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Only POST method supported"))
		case "POST":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil || len(body) < 1 {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Invalid payload"))
			} else {
				link := models.NewNextLink(string(body), crawler.Configuration.NumOfCrawls)
				crawler.Add(link)
				w.WriteHeader(http.StatusCreated)
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// stopAll calls Crawler.Stop which stops all crawling threads
func stopAll(crawler *crawler.Crawler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Stopping all threads")
		crawler.Stop()
		w.WriteHeader(http.StatusOK)
	}
}
