package main

import "fmt"

const helloPrefix = "Hello, "

func hello(name string) string {
	if name == "" {return helloPrefix + "World"}
	return helloPrefix + name
}

func main() {
	fmt.Println(hello("Tom"))
}
