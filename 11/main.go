// https://leetcode.com/problems/container-with-most-water/

package main

import (
	"fmt"
)

// go:inline
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// go:inline
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	maxArea := 0
	for j > i {
		maxArea = max((j-i)*min(height[i], height[j]), maxArea)
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
