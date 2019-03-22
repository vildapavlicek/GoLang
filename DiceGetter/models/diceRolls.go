package models

import (
	"net/http"

	dataparser "github.com/vildapavlicek/GoLang/DiceGetter/dataParser"
)

//DiceRolls main struct
type DiceRolls struct {
	NumOfRolls int
	Client     *http.Client
	Parser     dataparser.HTMLParser
}
