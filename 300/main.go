// https://leetcode.com/problems/longest-increasing-subsequence/

package main

import (
	"fmt"
	"sort"
)

func lengthOfLIS1(nums []int) int {
	r := make([]int, len(nums))
	for i := 0; i < len(r); i++ {
		r[i] = 1
	}
	max := 0
	for i := 0; i < len(r); i++ {
		lMax := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if r[j]+1 > lMax {
					lMax = r[j] + 1
				}
			}
		}
		r[i] = lMax
		if lMax > max {
			max = lMax
		}
	}
	return max
}

func lengthOfLIS(nums []int) int {
	numsSorted := make([]int, len(nums))
	k := 0
	for i := 0; i < len(nums); i++ {
		index := sort.SearchInts(numsSorted[0:k], nums[i])
		numsSorted[index] = nums[i]
		if index == k {
			k++
		}
	}
	return k
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 19}))
}
