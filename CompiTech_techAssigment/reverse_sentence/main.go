package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var data []string
	var n, i int
	var s string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Scanf("%d", &n)

	if n < 1 {
		fmt.Printf("Invalid input: Entered 0 or less lines to read or only string when integer was expected. Exiting program\n")
		os.Exit(1)
	}

	for scanner.Scan() {
		data = append(data, scanner.Text())
		if i == n {
			break
		}
		i++
	}

	for _, sentence := range data {
		s = ""
		words := strings.Split(sentence, " ")
		for i := len(words) - 1; i >= 0; i-- {
			s = s + words[i] + " "
		}
		fmt.Println(strings.Trim(s, " "))
	}
}
