package main

func main() {
	
}

func TwoSumBest(numbers []int, target int) [2]int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i+1; j < len(numbers); j++ {
			if numbers[i] + numbers[j] == target {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{0, 0}
}
