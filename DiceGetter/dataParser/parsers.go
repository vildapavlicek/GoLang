package dataparser

import (
	"net/http"
	"strconv"

	"golang.org/x/net/html"
)

//DiceRolls struct which holds parsed values
type DiceRolls struct {
	Data []int
}

//HTMLParser type aliasing for HTMLParsers
type HTMLParser func(response *http.Response) ([]int, error)

/*
//ParseHTML parser for results from page https://www.random.org returns array of results
func (d *DiceRolls) ParseHTML(response *http.Response) ([]int, error) {
	defer response.Body.Close()
	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, err
	}

	d.parseAltAtribute(doc)

	return d.Data, nil
}

func (d *DiceRolls) parseAltAtribute(n *html.Node) {
	//fmt.Printf("Parsing data.\n")
	if n.Data == "img" && n.Type == html.ElementNode {
		fmt.Printf("Found img tag!\n")
		d.Data = append(d.Data, d.parseRollNumber(n))
		fmt.Printf("Data values is: %v; ", d.Data)
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		d.parseAltAtribute(child)
	}
}

func (d *DiceRolls) parseRollNumber(n *html.Node) int {
	fmt.Println("Parsing roll number")
	var i int
	var err error

	for _, att := range n.Attr {
		if att.Key == "alt" {
			fmt.Printf("alt tag found with value: %v; ", att.Val)
			if i, err = strconv.Atoi(att.Val); err != nil {
				return -1
			}
			fmt.Printf("Returning value %v\n", i)
			return i
		}
	}
	return i
}
*/

//ParseHTML parser for results from page https://www.random.org returns array of results
func ParseHTML(response *http.Response) error {
	defer response.Body.Close()
	doc, err := html.Parse(response.Body)
	if err != nil {
		return err
	}
	parseAltAtribute(doc)
	return nil
}

func parseAltAtribute(n *html.Node) {
	//fmt.Printf("Parsing data.\n")
	if n.Data == "img" && n.Type == html.ElementNode {
		//store data
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		parseAltAtribute(child)
	}
}

func parseRollNumber(n *html.Node) int {
	var i int
	var err error

	for _, att := range n.Attr {
		if att.Key == "alt" {
			if i, err = strconv.Atoi(att.Val); err != nil {
				return -1
			}
			return i
		}
	}
	return i
}
