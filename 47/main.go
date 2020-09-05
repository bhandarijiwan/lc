// https://leetcode.com/problems/permutations-ii/

package main

import (
	"fmt"
)

func backTrack(nums []int, n, first int, output *[][]int) {
	if n == first {
		a := make([]int, len(nums))
		copy(a, nums)
		*output = append(*output, a)
	}
	m := make(map[int]bool)
	for i := first; i < len(nums); i++ {
		if i != first && nums[i] == nums[first] {
			continue
		}
		if _, ok := m[nums[i]]; ok {
			continue
		}
		nums[first], nums[i] = nums[i], nums[first]
		backTrack(nums, n, first+1, output)
		nums[first], nums[i] = nums[i], nums[first]
		m[nums[i]] = true

	}
}

func permuteUnique(nums []int) [][]int {
	output := make([][]int, 0)
	backTrack(nums, len(nums), 0, &output)
	return output
}

func main() {
	fmt.Println(permuteUnique([]int{2, 2, 2, 1, 1, 1}))
}
