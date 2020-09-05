// https://leetcode.com/problems/jump-game/
package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	l := len(nums)
	i := l - 1
	for j := i; j >= 0; j-- {
		if i <= j+nums[j] {
			i = j
		}
	}
	if i == 0 {
		return true
	}
	return false
}

func main() {
	// fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	// fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	// fmt.Println(canJump([]int{1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(canJump([]int{1, 1, 3, 0, 0, 1, 1, 1}))
	// fmt.Println(canJump([]int{0, 1}))
}
