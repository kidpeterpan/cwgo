package main

import (
	"fmt"
	"math"
)

// Complete the square sum function so that it squares each number passed into it and then sums the results together.
// For example, for [1, 2, 2] it should return 9 because 1^2 + 2^2 + 2^2 = 9.
func main() {
	input := []int{1, 2, 2}
	output := SquareSum(input)
	fmt.Println(output)
}

func SquareSum(numbers []int) int {
	total := 0
	for _,v := range numbers {
		s := math.Pow(float64(v),2)
		total += int(s)
	}
	return total
}