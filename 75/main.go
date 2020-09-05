// https://leetcode.com/problems/sort-colors/
// #counting sort,

package main

import (
	"fmt"
)

func sortColors1(nums []int) {
	c := make([]int, 3)
	for i := 0; i < len(nums); i++ {
		c[nums[i]]++
	}
	for i, j := 0, 0; i < len(nums); j = j + 1 {
		for k := 0; k < c[j]; k++ {
			nums[i] = j
			i++
		}
	}
}

func sortColors(nums []int) {
	i, j, k := 0, 0, len(nums)-1
	for k >= j {
		if nums[j] == 2 {
			nums[j], nums[k] = nums[k], nums[j]
			k--
		}
		if nums[j] <= 1 {
			if nums[j] == 0 {
				nums[j], nums[i] = nums[i], nums[j]
				i++
			}
			j++
		}
	}
}

func main() {
	input := []int{1, 0, 0, 0}
	sortColors(input)
	fmt.Println(input)
	input = []int{2, 0, 2, 1, 1, 0}
	sortColors(input)
	fmt.Println(input)
	input = []int{}
	sortColors(input)
	fmt.Println(input)
	input = []int{1, 1, 2}
	sortColors(input)
	fmt.Println(input)
	input = []int{2, 2, 1, 0, 2}
	sortColors(input)
	fmt.Println(input)
	input = []int{2, 2, 1, 0, 2}
	sortColors(input)
	fmt.Println(input)
}
