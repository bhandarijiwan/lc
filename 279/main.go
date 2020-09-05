// https://leetcode.com/problems/perfect-squares/
package main

import (
	"fmt"
	"math"
)

func backTrack(nums []int, inter []int, sum int, index int, count *int, nearest *int) {
	if sum < 0 {
		return
	}
	if sum == 0 {
		*count = len(inter)
		if *count < *nearest {
			*nearest = *count
		}
		return
	}

	for i := index; i < len(nums); i++ {
		inter = append(inter, nums[i])
		backTrack(nums, inter, sum-nums[i], index, count, nearest)
		inter = inter[:len(inter)-1]
	}
}

func numSquares(n int) int {
	if n < 4 {
		return n
	}
	s := int(math.Sqrt(float64(n)))
	nums := make([]int, s)
	for i, j := s, 0; i >= 1; i, j = i-1, j+1 {
		nums[j] = i * i
	}
	var inter []int
	minCount, count := math.MaxInt32, 0
	backTrack(nums, inter, n, 0, &count, &minCount)
	return minCount
}

func main() {
	// fmt.Println(numSquares(5))
	// fmt.Println(numSquares(6))
	// fmt.Println(numSquares(7))
	// fmt.Println(numSquares(10))
	// fmt.Println(numSquares(11))
	// fmt.Println(numSquares(14))
	// fmt.Println(numSquares(19))
	// fmt.Println(numSquares(50))
	// fmt.Println(numSquares(80))
	fmt.Println(numSquares(598))
}
