// https://leetcode.com/problems/missing-element-in-sorted-array/

package main

import (
	"fmt"
)

func missingElement(nums []int, k int) int {
	j := 0
	prev := nums[0] - 1
	for _, n := range nums {
		a := j + n - prev - 1
		if a >= k {
			return prev + k - j
		}
		j = a
		prev = n
	}
	return prev + k - j
}

func main() {
	fmt.Println(missingElement([]int{4, 7, 9, 10}, 1))
	fmt.Println(missingElement([]int{4, 7, 9, 10}, 3))
	fmt.Println(missingElement([]int{1, 2, 4}, 3))
	fmt.Println(missingElement([]int{1, 3}, 3))
}
