package main

import "fmt"

func main() {
	fmt.Println(findErrorNums([]int{2, 3, 2}))
}

// "1 1" -> "1 2"
// "2 2" -> "2,1"
// "1,2,2"-> "2,3"
// "1,3,3"-> "3,2"
// [3,2,2] -> "2,1"
// [2,3,2] ->"[2,1]"
//"[3,3,1] ->"[3,2]"

// if sorted - by index
func findErrorNumsInSorted(nums []int) []int {
	d := nums[0]
	m := 0
	searchDoneD := false
	searchDoneM := false
	for i := 1; i < len(nums); i++ {
		if !searchDoneD {
			for j := i; j < len(nums); j++ {
				if nums[j] == d {
					d = nums[j]
					searchDoneD = true
					break
				}
			}
		}

		if !searchDoneM {
			if nums[i-1] != i {
				m = i
				searchDoneM = true
			} else if nums[i] != i+1 {
				m = i + 1
				searchDoneM = true
			}
		}

		if searchDoneD && searchDoneM {
			break
		} else {
			d = nums[i]
		}

	}

	return []int{d, m}
}

// simple way with count
func findErrorNumsSimple(nums []int) []int {
	n := len(nums)
	count := make([]int, n+1)

	for _, num := range nums {
		count[num]++
	}

	duplicate, missing := 0, 0
	for i := 1; i <= n; i++ {
		if count[i] == 2 {
			duplicate = i
		} else if count[i] == 0 {
			missing = i
		}
	}

	return []int{duplicate, missing}
}

// math way
func findErrorNumsMathWay(nums []int) []int {
	n := len(nums)
	sum := 0
	sumOfSquares := 0
	expectedSum := n * (n + 1) / 2
	expectedSumOfSquares := n * (n + 1) * (2*n + 1) / 6

	for i := 0; i < n; i++ {
		sum += nums[i]
		sumOfSquares += nums[i] * nums[i]
	}

	// diff = duplicate - missing
	diff := sum - expectedSum

	// squareDiff = duplicate^2 - missing^2
	squareDiff := sumOfSquares - expectedSumOfSquares

	// duplicate + missing = squareDiff / diff
	sumDM := squareDiff / diff

	duplicate := (sumDM + diff) / 2
	missing := (sumDM - diff) / 2

	return []int{duplicate, missing}
}

// simple way with map
func findErrorNumsSimpleMath(nums []int) []int {
	seen := make(map[int]bool)
	duplicate := 0
	sum := 0
	n := len(nums)

	for _, num := range nums {
		if seen[num] {
			duplicate = num
		}
		seen[num] = true
		sum += num
	}

	expectedSum := n * (n + 1) / 2
	missing := expectedSum - (sum - duplicate)

	return []int{duplicate, missing}
}

// hard way with byte arephmetic
func findErrorNums(nums []int) []int {
	n := len(nums)
	xor := 0

	// 1. XOR all numbers in the array and numbers from 1 to n
	for i := 0; i < n; i++ {
		xor ^= nums[i] // числа из массива
		xor ^= (i + 1) // числа от 1 до n
	}

	// Now xor = duplicate XOR missing
	// because all other numbers have been canceled (a XOR a = 0)

	// 2. Find any set bit in xor
	// This is the bit where duplicate and missing differ
	rightmostBit := xor & -xor

	// 3. Divide the numbers into two groups by this bit
	duplicate, missing := 0, 0

	for i := 0; i < n; i++ {
		// Group the numbers from the array
		if nums[i]&rightmostBit != 0 {
			duplicate ^= nums[i]
		} else {
			missing ^= nums[i]
		}

		// Group numbers from 1 to n
		if (i+1)&rightmostBit != 0 {
			duplicate ^= (i + 1)
		} else {
			missing ^= (i + 1)
		}
	}

	// 4. Determine which of the numbers is duplicate and which is missing
	for i := 0; i < n; i++ {
		if nums[i] == duplicate {
			return []int{duplicate, missing}
		}
	}

	return []int{missing, duplicate}
}
