// https://leetcode.com/problems/range-sum-query-mutable/

package main

import (
	"fmt"
)

type NumArray struct {
	d []int
	n int
}

func Constructor(nums []int) NumArray {
	st := NumArray{n: len(nums)}
	st.build(nums)
	return st
}

func (na *NumArray) build(nums []int) {
	l := len(nums) << 1
	na.d = make([]int, l)
	copy(na.d[len(nums):], nums)
	for i := na.n - 1; i > 0; i-- {
		na.d[i] = na.d[i*2] + na.d[(i*2)+1]
	}
}

func (na *NumArray) Update(i int, val int) {
	j := i + na.n
	na.d[j] = val
	for j > 1 {
		j >>= 1
		na.d[j] = na.d[j<<1] + na.d[(j<<1)+1]
	}
}

func (na *NumArray) SumRange(i int, j int) int {
	left, right := i+na.n, j+na.n+1
	result := 0
	for left < right {
		if left&1 == 1 {
			result += na.d[left]
			left++
		}
		if right&1 == 1 {
			right--
			result += na.d[right]
		}
		left = left >> 1
		right = right >> 1
	}
	return result
}

func main() {
	obj := Constructor([]int{1, 3, 5, 9, 8, 4, 8, 0, 4, 1, 4, 3, 9, 7})
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(0, 3))
	fmt.Println(obj.SumRange(0, 1))
	fmt.Println(obj.SumRange(0, 0))
	fmt.Println(obj.SumRange(0, 13))
	obj.Update(12, 20)
	fmt.Println(obj.SumRange(0, 13))
	fmt.Println(obj.SumRange(12, 13))
	// fmt.Println(obj.SumRange(0, 3))
	// fmt.Println(obj.SumRange(0, 2))
	// fmt.Println(obj.SumRange(1, 2))
}
