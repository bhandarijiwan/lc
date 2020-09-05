// https://leetcode.com/problems/search-in-rotated-sorted-array/

package main

import (
	"fmt"
	"sort"
)

func search(nums []int, target int) int {
	// Figure out the pivot
	l := len(nums)
	if l < 1 {
		return -1
	}
	i, j := 0, l-1
	for nums[i] > nums[j] && j-i > 1 {
		m := i + (j-i)/2
		if nums[m] > nums[i] {
			i = m + 1
		} else {
			j = m
		}
	}
	if nums[j] < nums[i] {
		i = j
	}
	index := sort.Search(l, func(idx int) bool {
		r := (i + idx) % l
		return nums[r] >= target
	})
	if index < l && nums[(i+index)%l] == target {
		return (i + index) % l
	}
	return -1
}

func main() {
	fmt.Println(search([]int{7, 0, 1, 2, 4, 5, 6}, 0))
	fmt.Println(search([]int{6, 7, 0, 1, 2, 4, 5}, 0))
	fmt.Println(search([]int{5, 6, 7, 0, 1, 2, 4}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{2, 4, 5, 6, 7, 0, 1}, 0))
	fmt.Println(search([]int{1, 2, 4, 5, 6, 7, 0}, 0))
	fmt.Println(search([]int{0, 1, 2, 4, 5, 6, 7}, 0))

}
