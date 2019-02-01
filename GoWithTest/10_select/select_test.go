package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("basic function test", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("Got %v; want %v", got, want)
		}
	})

	t.Run("timeout test", func(t *testing.T) {
		serverA := makeDelayedServer(2 * time.Second)
		defer serverA.Close()
		serverB := makeDelayedServer(2 * time.Second)
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 1 * time.Second)

		if err == nil {
			t.Fatal("error expected but got none")
		}

	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
