package validator

import (
	"sync"

	"github.com/vildapavlicek/GoLang/crawler/pkg/models"
)

type Validate func(data models.DataHolder, out chan models.DataHolder)

func Execute(ID string, input <-chan models.DataHolder, validate Validate) <-chan models.DataHolder {
	var wg sync.WaitGroup
	wg.Add(1)

	out := make(chan models.DataHolder)

	go func() {
		for data := range input {
			validate(data, out)
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
