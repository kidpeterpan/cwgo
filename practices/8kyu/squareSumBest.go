package main

func main() {
	
}

func SquareSumBest(nums []int) (res int) {
	for _, val := range nums {
		res += val * val
	}
	return res
}
