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
func SetHandlers(m *http.ServeMux, c *crawler.Crawler, s *http.Server) {
	m.HandleFunc("/", index)
	m.HandleFunc("/api/v1/link", linkHandler(c, handler))
	m.HandleFunc("/api/v1/stop", stopAll(c,s, handler))
}

//TODO should be used for landing page, so far used for testing tamplates
func index(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./views/index.gohtml"))
	tpl.Execute(w, "Vilda")
}

// empty handler used for wrapping
func handler(w http.ResponseWriter, r *http.Request) {
}

// accepts POST method to add new link for crawling if successful returns StatusCreated - 201 else StatusBadRequest 400
// GET method returns http.StatusMethodNotAllowed - 405
// default response set to http.StatusInternalServerError - 500
func linkHandler(crawler *crawler.Crawler, h http.HandlerFunc) http.HandlerFunc {
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
func stopAll(crawler *crawler.Crawler,s *http.Server, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Stopping all threads")
		crawler.Stop()
		w.WriteHeader(http.StatusOK)
	}
}
