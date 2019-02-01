package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var m = make(map[string]int)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		line := strings.Split(sc.Text(), " ")
		switch {
		case line[0] == "COPY":
			copy(line)
		case line[0] == "ADD":
			add(line)
		case line[0] == "PRINT":
			print(line)
		case isComment(line):
			break
		default:
			fmt.Println("Error!")
		}
	}
}

func isComment(data []string) bool {
	if len(data) <= 1 {
		return false
	}

	s := data[0]
	if r := s[0]; r == '#' {
		return true
	}
	return false
}

func copy(data []string) {
	x := data[1]
	y := data[2]

	if !strings.HasPrefix(data[1], "_") {
		fmt.Println("Error!")
		return
	}

	if i, b := isInt(y); b {
		m[x] = i
	} else if strings.HasPrefix(x, "_") {
		i, b := m[y]
		if b {
			m[x] = i
		} else {
			fmt.Println("Error!")
		}
	}

}

func add(data []string) {
	x := data[1]
	y := data[2]

	if i, b := isInt(y); b {
		m[x] = m[x] + i
	} else if i, b := m[y]; b {
		m[x] = m[x] + i
	} else {
		fmt.Println("Error!")
	}
}

func print(data []string) {
	v := data[1]
	toPrint, exists := m[v]
	if exists {
		fmt.Println(toPrint)
	} else {
		fmt.Println("Error!")
	}
}

func isInt(s string) (int, bool) {
	i, err := strconv.Atoi(s)

	if err == nil {
		return i, true
	} else {
		return 0, false
	}
}
