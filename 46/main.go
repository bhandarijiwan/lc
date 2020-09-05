// https://leetcode.com/problems/permutations/

package main

import (
	"fmt"
)

func permute1(nums []int) [][]int {
	l := len(nums)
	if l == 0 {
		return [][]int{}
	}
	if l == 1 {
		return [][]int{nums}
	}
	r := make([][]int, 0)
	for i := 0; i < l; i++ {
		subarray := make([]int, 0)
		subarray = append(subarray, nums[0:i]...)
		subarray = append(subarray, nums[i+1:l]...)
		pSubArray := permute(subarray)
		for _, ps := range pSubArray {
			a := []int{nums[i]}
			a = append(a, ps...)
			r = append(r, a)
		}
	}
	return r
}

func backTrack(nums []int, n, first int, output *[][]int) {
	if first == n {
		a := make([]int, len(nums))
		copy(a, nums)
		*output = append(*output, a)
	}
	for i := first; i < len(nums); i++ {
		nums[first], nums[i] = nums[i], nums[first]
		backTrack(nums, n, first+1, output)
		nums[first], nums[i] = nums[i], nums[first]
	}
}

func permute(nums []int) [][]int {
	output := make([][]int, 0)
	backTrack(nums, len(nums), 0, &output)
	return output
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
