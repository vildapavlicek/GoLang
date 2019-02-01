package main

import (
	"errors"
	"net/http"
	"time"
)

var ErrTimeOut = errors.New("request timed out")
var tenSecondTimeout = 10 * time.Second

func main() {

}

func Racer(firstUrl, secondUrl string) (winner string, err error) {
	return ConfigurableRacer(firstUrl, secondUrl, tenSecondTimeout)
}

func ConfigurableRacer(firstUrl, secondUrl string, timeout time.Duration) (winner string, err error) {

	select {
	case <-ping(firstUrl):
		return firstUrl, nil
	case <-ping(secondUrl):
		return secondUrl, nil
	case <-time.After(timeout):
		return "", ErrTimeOut
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)

	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
