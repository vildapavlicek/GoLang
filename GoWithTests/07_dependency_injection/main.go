package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
Greet("Elodie", os.Stdout)
}

func Greet(name string, output io.Writer)  {
	fmt.Fprintf(output, "Hello, %s", name)
}