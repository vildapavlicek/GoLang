package main

import (
	"fmt"
)

func main() {
	var data [3]int
start:
	data = getInput()

	if !isValid(data) {
		fmt.Printf("Invalid input. Expected 3 numbers in range 1 to 9, got %v\n", data)
		goto start
	}

	fmt.Println(fibonacci(data))
}

func getInput() [3]int {
	var a int
	var read [3]int
	fmt.Print("Enter 3 numbers where 0 < n < 10 and hit Enter: ")

	for i := 0; i < 3; i++ {
		fmt.Scanf("%d",&a)
		read[i] = a
	}
	return read
}

func isValid(data [3]int) bool {
	for _, n := range data {
		if n < 1 || n > 10 {
			return false
		}
	}

	return true
}

func fibonacci(data [3]int) int {
	nums := data[0:2]
	index := data[2]
	for i := 0; i < index; i++ {
		nums = append(nums, nums[i] + nums[i+1])
	}

	return nums[index-1]
}