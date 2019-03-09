package crawler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetResponse(t *testing.T) {
	t.Run("OK Response", func(t *testing.T) {
		status := http.StatusOK
		want := http.StatusOK
		server := makeHttpServer(status)
		defer server.Close()
		got := getResponse("GET", server.URL, "", myClient)
		defer got.Body.Close()
		assertStatusEquals(t, want, got.StatusCode)
	})
}

func TestParseNextVideoData(t *testing.T) {

	t.Run("ParseNextVideoData from response.dat", func(t *testing.T) {
		file, err := os.Open("response.dat")
		if err != nil {
			t.Errorf("Failed to open file with test data; reason: %s", err)
		}
		body, err := ioutil.ReadAll(file)
		if err != nil {
			t.Errorf("Failed to read test data from file; reason: %s", err)
		}
		server := makeFakeYoutubeServer(body)
		defer server.Close()

		res := getResponse("GET", server.URL, "", myClient)
		defer res.Body.Close()
		wantLink := "/watch?v=KR-eV7fHNbM"
		wantTitle := "TheFatRat - The Calling (feat. Laura Brehm)"
		gotLink, gotTitle, err := parseNextVideoData(res)
		if err != nil {
			t.Errorf("Failed to parse response body; %s; reason: %s", res.Body, err)
		}

		assertLinkEquals(t, wantLink, gotLink)
		assertTitleEquals(t, wantTitle, gotTitle)

	})

	t.Run("ParseNextVideoData from response2.dat", func(t *testing.T) {
		file, err := os.Open("response2.dat")
		if err != nil {
			t.Errorf("Failed to open file with test data; reason: %s", err)
		}
		body, err := ioutil.ReadAll(file)
		if err != nil {
			t.Errorf("Failed to read test data from file; reason: %s", err)
		}
		server := makeFakeYoutubeServer(body)
		defer server.Close()

		res := getResponse("GET", server.URL, "", myClient)
		defer res.Body.Close()
		wantLink := "/watch?v=TsTFVdcpLrE"
		wantTitle := "Hans Zimmer - Time ( Cyberdesign Remix )"
		gotLink, gotTitle, err := parseNextVideoData(res)
		if err != nil {
			t.Errorf("Failed to parse response body; %s; reason: %s", res.Body, err)
		}

		assertLinkEquals(t, wantLink, gotLink)
		assertTitleEquals(t, wantTitle, gotTitle)
	})

}

func makeHttpServer(status int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
	}))
}

func makeFakeYoutubeServer(body []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(body)

	}))
}

func assertStatusEquals(t *testing.T, want, got int) {
	if want != got {
		t.Errorf("Got '%v', want: '%v'", got, want)
	}

}

func assertLinkEquals(t *testing.T, wantLink, gotLink string) {
	if wantLink != gotLink {
		t.Errorf("Got %s, want %s", gotLink, wantLink)
	}
}

func assertTitleEquals(t *testing.T, wantTitle, gotTitle string) {
	if wantTitle != gotTitle {
		t.Errorf("Got %s, want %s", gotTitle, wantTitle)
	}
}
