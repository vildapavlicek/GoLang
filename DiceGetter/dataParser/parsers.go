package dataParser

import "net/http"

type HtmlParser func(response *http.Response) ([]int, error)

func DiceRollsParser(response *http.Response) ([]int, error) {

	return []int{}, nil
}
