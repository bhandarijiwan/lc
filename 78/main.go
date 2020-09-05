// https://leetcode.com/problems/subsets/

package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	m := make(map[int]int)
	a := make([]int, 0)
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			a = append(a, num)
		}
		m[num] = num
	}
	nums = a
	output := make([][]int, len(nums))
	// size of 1
	sizeOffset := make([]int, len(nums))
	for index, ants := range nums {
		output[index] = []int{ants}
		sizeOffset[index] = index + 1
	}
	sizeBegin := 0
	for size := 2; size <= len(nums); size++ {
		l := len(output)
		// for each number
		count := 0
		for i := 0; i < len(nums); i++ {
			// for each subset of length size-1
			for j := sizeBegin + sizeOffset[i]; j < l; j++ {
				arr := make([]int, size)
				arr[0] = nums[i]
				copy(arr[1:], output[j])
				output = append(output, arr)
				count++
			}
			sizeOffset[i] = count
		}
		sizeBegin = l
	}
	output = append(output, []int{})
	return output
}

func main() {
	fmt.Println(subsets([]int{1, 2}))
}
