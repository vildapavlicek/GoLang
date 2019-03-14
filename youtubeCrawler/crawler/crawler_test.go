package crawler

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
	"youtubeCrawler/models"
	"youtubeCrawler/store"
)

type countParser struct {
}

func (cp countParser) ParseData(response *http.Response) (link, title string, err error) {
	return "", "", nil
}

type fakeStore struct {
	data []models.NextLink
}

func (fs fakeStore) Store(link models.NextLink) error {
	fs.data = append(fs.data, link)
	return nil
}

func (fs fakeStore) Close() {

}

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

func TestCrawl(t *testing.T) {
	t.Run("Test 30 iterations", func(t *testing.T) {
		testStore := fakeStore{
			data: make([]models.NextLink, 30),
		}
		testStoreManager := &store.Manager{
			StorePipe:        make(chan models.NextLink, 10),
			StoreDestination: testStore,
			Shutdown:         make(chan bool, 1),
		}

		server := makeHttpServer(200)
		defer server.Close()

		firstLink := models.NextLink{
			BaseUrl:       server.URL,
			Link:          "",
			NOfIterations: 30,
			Number:        0,
		}

		cp := countParser{}
				crawler := Crawler{
			data: make(chan models.NextLink, 5),
			parser: cp,
			wg: sync.WaitGroup{},
			stopSignal: make(chan bool),
			StoreManager: testStoreManager,
		}
		crawler.Add(firstLink)
		go crawler.crawl(1, crawler.parser)
		go crawler.StoreManager.StoreData()
		time.Sleep(3 * time.Second)
		crawler.Stop()

		wantLength := 30
		gotLength := len(testStore.data)

		assertLengthEquals(t, wantLength, gotLength)

	})

	t.Run("Test 20 iterations", func(t *testing.T) {
		testStore := fakeStore{
			data: make([]models.NextLink, 30),
		}
		testStoreManager := &store.Manager{
			StorePipe:        make(chan models.NextLink, 10),
			StoreDestination: testStore,
			Shutdown:         make(chan bool, 1),
		}

		server := makeHttpServer(200)
		defer server.Close()

		firstLink := models.NextLink{
			BaseUrl:       server.URL,
			Link:          "",
			NOfIterations: 20,
			Number:        0,
		}

		cp := countParser{}
		crawler := Crawler{
			data: make(chan models.NextLink, 5),
			parser: cp,
			wg: sync.WaitGroup{},
			stopSignal: make(chan bool),
			StoreManager: testStoreManager,
		}
		crawler.Add(firstLink)
		go crawler.crawl(1, crawler.parser)
		go crawler.StoreManager.StoreData()
		time.Sleep(3 * time.Second)
		crawler.Stop()

		wantLength := 30
		gotLength := len(testStore.data)

		assertLengthEquals(t, wantLength, gotLength)

	})
}



func makeHttpServer(status int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)

	}))
}

func assertStatusEquals(t *testing.T, want, got int) {
	t.Helper()
	if want != got {
		t.Errorf("Got '%v', want: '%v'", got, want)
	}

}

func assertLengthEquals(t *testing.T, want, got int) {
	t.Helper()
	if want != got {
		t.Errorf("Got '%v', want: '%v'", got, want)
	}
}
