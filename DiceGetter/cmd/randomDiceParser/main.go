package main

import (
	"fmt"
	"time"

	dataparser "github.com/vildapavlicek/GoLang/DiceGetter/dataParser"

	"github.com/vildapavlicek/GoLang/DiceGetter/httpClient"
)

func main() {
	timeout := 5 * time.Second
	client := httpClient.New(timeout)
	var data []int

	response, err := client.GetResponse("GET", "https://www.random.org/", "dice/?num=10", nil)
	if err != nil {
		fmt.Printf("Failed to get response; %s", err)
	}

	data, err = dataparser.ParseHTML(response, data)
	if err != nil {
		fmt.Printf("Data parsing failed; %s", err)
	}

	fmt.Println(data)
}

/*
Question 1: Write a program to do the following:
● TODO: Make an HTTP request to https://www.random.org/dice/?num=10 . Make sure to handle error statuses, timeouts, etc.
● TODO: Parse out the 10 dice values returned. For this example, let’s suppose that the values returned are as follows: 3,5,1,2,6,5,1,6,4,2
● TODO: Bucket the die’s by their value, and output the counts to stdout. Continuing with the aforementioned example data, the output ( format: die value -> count ) would look like:
○ 1 -> 2
○ 2 -> 2
○ 3 -> 1
○ 4 -> 1
○ 5 -> 2
○ 6 -> 2
● TODO: Sort the die rolls in increasing value order, and send the results to stderr. Continuing with the aforementioned example data, the output would look like:
○ 1 1 2 2 3 4 4 5 5 6 6
● TODO: Now, convert the sorted rolls to the following json format: { "dice": [ 1,1,2,2,3,4,5,5,6,6 ] }
● TODO: Make an HTTP POST with the json as the POST payload. requestbin (https://requestb.in/ ) is a nice free online utility that will allow you to make http requests.Make sure to handle error statuses, timeouts, etc.
*/
