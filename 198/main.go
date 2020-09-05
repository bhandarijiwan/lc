//https://leetcode.com/problems/house-robber/

package main

import (
	"fmt"
)

func max(n ...int) int {
	m := 0
	for _, n := range n {
		if n > m {
			m = n
		}
	}
	return m
}
func rob(nums []int) int {
	l := len(nums)
	if l <= 2 {
		if l > 0 {
			return max(nums...)
		}
		return 0
	}
	p2 := 0
	p1 := 0
	for i := l - 1; i >= 0; i-- {
		a := max(nums[i]+p2, p1)
		p2 = p1
		p1 = a
	}
	return p1
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
