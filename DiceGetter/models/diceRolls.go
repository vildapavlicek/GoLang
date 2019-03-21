package models

import (
	"../dataParser"
	"net/http"
)

type DiceRolls struct {
	NumOfRolls int
	Client     *http.Client
	Parser     dataParser.HtmlParser
	Data       []int `json:"data"`
}
