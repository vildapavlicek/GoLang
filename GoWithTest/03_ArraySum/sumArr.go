package main

import "fmt"

func main() {

	fmt.Println(SumAll([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}))
}

func Sum(arr []int) int {
	var result int

	for _, n := range arr {
		result += n
	}

	return result
}

func SumAll(data ...[]int) []int {
	var sums []int

	for _, slice := range data {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(data ...[]int) []int {
	var sums []int
	for _, slice := range data {
		if len(slice) < 1 {
			sums = append(sums, 0)
		} else {
			tails := slice[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
