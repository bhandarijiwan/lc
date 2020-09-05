// https://leetcode.com/problems/maximum-product-subarray/

package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	p, q := 1, 1
	c, r := 0, 0
	m := math.MinInt64
	x := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if m > r {
				r = m
			}
			c += 1
			x = make([]int, 0)
			m, p, q = math.MinInt64, 1, 1
		} else if nums[i] < 0 {
			q = q * nums[i]
			x = append(x, q)
			j := 0
			if len(x)%2 > 0 && len(x) > 1 {
				j = 1
			}
			s := x[j]
			for j = j + 1; j < len(x); j++ {
				s *= x[j]
			}
			if s > m {
				m = s
			}
			if s > 0 {
				p = s
			} else {
				p = 1
			}
			q = 1
		} else {
			q = q * nums[i]
			p = p * nums[i]
			if p > m {
				m = p
			}
		}
	}
	if c > 0 && r > m {
		return r
	}
	return m
}

func main() {
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{4, 4, -2, 12, 1}))
	fmt.Println(maxProduct([]int{4, 4, -1}))
	fmt.Println(maxProduct([]int{-2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 3, -2, 4, -3}))
	fmt.Println(maxProduct([]int{-2, -3, -2, -4, -3}))
	fmt.Println(maxProduct([]int{2}))
	fmt.Println(maxProduct([]int{-2}))
	fmt.Println(maxProduct([]int{-2, 0, -6, 2, -2, -2}))
	fmt.Println(maxProduct([]int{}))
}
