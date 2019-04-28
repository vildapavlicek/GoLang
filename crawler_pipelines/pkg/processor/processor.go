package processor

import (
	"sync"

	"github.com/vildapavlicek/GoLang/crawler/pkg/models"
)

type Process func(name string, data models.DataHolder, out chan models.DataHolder)

func Execute(ID string, process Process, inData <-chan models.DataHolder) <-chan models.DataHolder {
	var wg sync.WaitGroup

	out := make(chan models.DataHolder)

	go func() {
		for data := range inData {
			process(ID, data, out)
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
