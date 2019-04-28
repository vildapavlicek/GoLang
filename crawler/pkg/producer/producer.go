package producer

import (
	"sync"

	"github.com/vildapavlicek/GoLang/crawler/pkg/models"
)

type Producer func(ID string, out chan models.DataHolder)

func Execute(ID string, produce Producer) <-chan models.DataHolder {
	var wg sync.WaitGroup
	wg.Add(1)

	out := make(chan models.DataHolder)

	go func() {
		produce(ID, out)
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
