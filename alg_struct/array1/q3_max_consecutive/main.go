package main

import "fmt"

func main() {
	fmt.Println(findMaxConsecutiveOnes2([]int{1, 1, 0, 1, 1, 1}))
}

func findMaxConsecutiveOnes2(nums []int) int {
	max, curr := 0, 0
	for _, n := range nums {
		if n == 1 {
			curr++
			if curr > max {
				max = curr
			}
		} else {
			curr = 0
		}
	}
	return max

}
