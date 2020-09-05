// https://leetcode.com/problems/majority-element/

package main

import (
	"fmt"
)

func majorityElement1(nums []int) int {
	m := make(map[int]int)
	majorElement, maxCount := 0, len(nums)/2
	for _, num := range nums {
		m[num]++
		if m[num] > maxCount {
			majorElement = num
			break
		}
	}
	return majorElement
}

// Boyer-Moore majority vote algorithm
func majorityElement(nums []int) int {
	majority, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			majority = num
		}
		if num == majority {
			count++
		} else {
			count--
		}
	}
	return majority
}

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElement([]int{3, 3, 4}))
}
