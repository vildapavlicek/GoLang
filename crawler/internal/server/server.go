package server

import (
	"fmt"
	"net/http"
	"time"
)

func newServer(address string, mux *http.ServeMux) *http.Server {

	return &http.Server{
		Addr:              address,
		ReadHeaderTimeout: 20000 * time.Millisecond,
		ReadTimeout:       20000 * time.Millisecond,
		WriteTimeout:      20000 * time.Millisecond,
		IdleTimeout:       60000 * time.Millisecond,
		Handler:           mux,
	}

}

func newMuxWithHandlers() *http.ServeMux {
	m := http.NewServeMux()
	registerHandlers(m)
	return m

}

func registerHandlers(m *http.ServeMux) {
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
