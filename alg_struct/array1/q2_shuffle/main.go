package main

import "fmt"

func main() {
	fmt.Println(shuffle([]int{1, 2, 3, 4, 5, 6}, 3))
}

func shuffle(nums []int, n int) []int {
	a1 := append(make([]int, 0, n), nums[:n]...)
	a2 := append(make([]int, 0, n), nums[n:]...)
	ans := make([]int, 0, 2*n)
	for i := 0; i < n; i = i + 1 {
		ans = append(ans, a1[i])
		ans = append(ans, a2[i])
	}

	return ans
}

func shuffle2(nums []int, n int) []int {
	ans := make([]int, 2*n)
	for i := 0; i < n; i++ {
		ans[2*i] = nums[i]
		ans[2*i+1] = nums[i+n]
	}
	return ans
}
