package server

import (
	"fmt"

	"github.com/gorilla/mux"

	"net/http"
	"time"
)

func newServer(address string, m *mux.Router) *http.Server {

	return &http.Server{
		Addr:              address,
		ReadHeaderTimeout: 20000 * time.Millisecond,
		ReadTimeout:       20000 * time.Millisecond,
		WriteTimeout:      20000 * time.Millisecond,
		IdleTimeout:       60000 * time.Millisecond,
		Handler:           m,
	}
}

func newMuxWithHandlers() *mux.Router {
	m := mux.NewRouter()

	registerHandlers(m)
	return m

}

func registerHandlers(m *mux.Router) {
	//m.Handle("")
}

func Start() {
	m := newMuxWithHandlers()
	server := newServer(":8080", m)

	err := server.ListenAndServe()

	if err != nil {
		fmt.Printf("Failed to start http.Server, err: '%s'", err)
	}
}
