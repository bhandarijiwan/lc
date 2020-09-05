// https://leetcode.com/problems/product-of-array-except-self/

package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	p := 1
	r := make([]int, len(nums))
	r[0] = 1
	for i := 1; i < len(nums); i++ {
		p = p * nums[i-1]
		r[i] = p
	}
	p = 1
	for j := len(nums) - 2; j >= 0; j-- {
		p = p * nums[j+1]
		r[j] = r[j] * p
	}
	return r
}

func main() {
	fmt.Println(productExceptSelf([]int{4, 3, 2, 1, 2}))
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4, 5}))
}
