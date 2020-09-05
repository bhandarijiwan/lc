// https://leetcode.com/problems/shortest-unsorted-continuous-subarray/

package main

import (
	"fmt"
	"math"
)

func minmax(nums []int) (int, int) {
	min, max := math.MaxInt32, math.MinInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min, max
}

func findUnsortedSubarray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	min, max := minmax(nums)
	i, j := 0, len(nums)
	if nums[i] == min {
		for i = 1; i < len(nums); i++ {
			if nums[i] < min {
				i--
				break
			}
			min = nums[i]
		}
	}
	if i >= len(nums) {
		return 0
	}

	if nums[j-1] == max {
		for j = j - 1; j >= 0; j-- {
			if nums[j] > max {
				j = j + 2
				break
			}
			max = nums[j]
		}
	}
	min, max = minmax(nums[i:j])
	for ; i >= 0 && nums[i] > min; i-- {
	}
	i++
	for ; j < len(nums) && nums[j] < max; j++ {
	}
	return j - i
}

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}) == 5)
	fmt.Println(findUnsortedSubarray([]int{2, 4, 3}) == 2)

	fmt.Println(findUnsortedSubarray([]int{2}) == 0)
	fmt.Println(findUnsortedSubarray([]int{2, 3, 0}) == 3)
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 2, 3, 0}) == 6)
	fmt.Println(findUnsortedSubarray([]int{2, 3, 3, 1, 4, 5}) == 4)
	fmt.Println(findUnsortedSubarray([]int{}) == 0)
	fmt.Println(findUnsortedSubarray([]int{1, 1, 1, 1}) == 0)
	fmt.Println(findUnsortedSubarray([]int{2, 3, 3, 2, 4}) == 3)
	fmt.Println(findUnsortedSubarray([]int{3, 2, 1, 0}) == 4)
	fmt.Println(findUnsortedSubarray([]int{2, 2, 3, 5, 2, 2, 3, 3, 4, 4, 5, 8}) == 8)
	fmt.Println(findUnsortedSubarray([]int{1, 2, 4, 5, 3}) == 3)
	fmt.Println(findUnsortedSubarray([]int{1, 3, 5, 4, 2}) == 4)
	fmt.Println(findUnsortedSubarray([]int{1, 3, 5, 4, 5, 5, 5}) == 2)
	fmt.Println(findUnsortedSubarray([]int{1, 4, 5, 3}) == 3)
	fmt.Println(findUnsortedSubarray([]int{2, 1, 3, 4, 5}) == 2)
	fmt.Println(findUnsortedSubarray([]int{2, 1, 3, 4, 5}) == 2)
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4, 5, 65, 45, 23, 23, 23, 44, 100, 101, 102, 103, 342, 512, 312, 423, 454, 623, 231, 111, 235, 1001, 1002, 1003, 1004}) == 19)
}
