//https://leetcode.com/problems/maximum-subarray/

package main

import "fmt"

func maxSubArray(nums []int) int {
	l := len(nums)
	m := nums[l-1]
	p := m
	for i := l - 2; i >= 0; i-- {
		a := nums[i]
		p = p + a
		if a > p {
			p = a
		}
		if p > m {
			m = p
		}
	}
	return m
}

func max(nums ...int) int {
	l := nums[0]
	for _, a := range nums {
		if a > l {
			l = a
		}
	}
	return l
}

func maxSubArray2(nums []int) int {
	l := len(nums)
	if l == 0 {
		return l
	}
	if l == 1 {
		return nums[0]
	}
	m := l / 2
	left := maxSubArray2(nums[0:m])
	right := maxSubArray(nums[m:l])
	s := left + right
	return max(left, right, s)
}

func main() {
	fmt.Println(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
