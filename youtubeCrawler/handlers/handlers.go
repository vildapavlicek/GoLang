package handlers

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"youtubeCrawler/crawler"
	"youtubeCrawler/models"
)

func SetHandlers(m *http.ServeMux, c *crawler.Crawler) {
	m.HandleFunc("/", index)
	m.HandleFunc("/api/v1/link", linkHandler(c, handler))
	m.HandleFunc("/api/v1/stop", stopAll(c, handler))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./views/index.gohtml"))
	tpl.Execute(w, "Vilda")
}

func handler(w http.ResponseWriter, r *http.Request) {
}

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
				link := models.NewNextLink(string(body), 0)
				crawler.Add(link)
				w.WriteHeader(http.StatusCreated)
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func stopAll(crawler *crawler.Crawler, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		crawler.Stop()
		w.WriteHeader(http.StatusOK)
	}
}
