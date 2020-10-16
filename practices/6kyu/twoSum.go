package main

import "fmt"

/*
Write a function that takes an array of numbers (integers for the tests) and a target number.

- It should find two different items in the array that, when added together,
- give the target value. The indices of these items should then be returned in a tuple like so: (index1, index2).
For the purposes of this kata, some tests may have multiple answers; any valid solutions will be accepted.

The input will always be valid (numbers will be an array of length 2 or greater, and all of the items will be numbers;
target will always be the sum of two different items from that array).

Based on: http://oj.leetcode.com/problems/two-sum/

twoSum [1, 2, 3] 4 === (0, 2)
*/

func main() {
	output := TwoSum([]int{1,2,3}, 4)
	fmt.Println(output)
}

func TwoSum(numbers []int, target int) [2]int {
	for i,v1 := range numbers {
		for j,v2 := range numbers {
			if i == j {
				continue
			}
			if (v1 + v2) == target {
				return [2]int{i,j}
			}
		}
	}

	return [2]int{0,0}
}