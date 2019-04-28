package youtube

import (
	"os"
	//"sync"

	"github.com/vildapavlicek/GoLang/crawler/internal/server"
	youtubeclient "github.com/vildapavlicek/GoLang/crawler/internal/youtubeClient"
	"github.com/vildapavlicek/GoLang/crawler/pkg/models"
)

type youtubeCrawler struct {
	client *youtubeclient.Client
	ID     string
}

func (y *youtubeCrawler) Crawl(ID string, out chan models.Data) {
	//var wg sync.WaitGroup
	var incomingRequests = make(chan string)
	var shutdown = make(chan os.Signal, 1)

	y.initClient()

	go server.Start()
	//wg.Add(1)

	for {
		select {
		case <-incomingRequests:

		case <-shutdown:
		default:
		}
	}

	//wg.Wait()
}

func (y *youtubeCrawler) ParseLink(suffix string) {
	res, err := y.client.GetResponse(suffix)
	if err != nil {
		//check error
	}

	title, link, err := Parse(res)
	if err != nil {
		//check err
	}

}

func (y *youtubeCrawler) initClient() {
	y.client = youtubeclient.New(30, 10, 60)
}
