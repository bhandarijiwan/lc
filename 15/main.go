// https://leetcode.com/problems/3sum/
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	l := len(nums)
	if l < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	output := make([][]int, 0)
	for i := 0; i < l; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, l-1; j < k; {
			if j > i+1 && nums[j-1] == nums[j] {
				j += 1
				continue
			}
			if k < l-1 && nums[k] == nums[k+1] {
				k -= 1
				continue
			}
			s := nums[i] + nums[j] + nums[k]
			if s > 0 {
				k = k - 1
			} else if s < 0 {
				j = j + 1
			} else {
				output = append(output, []int{nums[i], nums[j], nums[k]})
				j, k = j+1, k-1
			}
		}
	}
	return output
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0, 0, 0, 0, 0, 1}))
	fmt.Println(threeSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}))

}
