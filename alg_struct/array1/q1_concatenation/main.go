package main

import "fmt"

func getConcatenation(nums []int) []int {
	ans := append(append(make([]int, 0, len(nums)), nums...), nums...)

	return ans
}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 3}))
}
