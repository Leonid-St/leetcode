package main

import "fmt"

func main() {
	fmt.Println(smallerNumbersThanCurrent1([]int{8, 1, 2, 2, 3}))
}

// one if
func smallerNumbersThanCurrent1(nums []int) []int {
	m := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if (j != i) && (nums[j] < nums[i]) {
				m[i]++
			}
		}

	}

	return m
}

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if nums[j] < nums[i] {
				count++
			}
		}
		result[i] = count
	}

	return result
}

func smallerNumbersThanCurrentFix(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	index := make([]int, 101)

	for i := 0; i < n; i++ {
		index[nums[i]]++
	}

	for i := 1; i <= 100; i++ {
		index[i] += index[i-1]
	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			result[i] = 0
		} else {
			result[i] = index[nums[i]-1]
		}
	}

	return result
}
