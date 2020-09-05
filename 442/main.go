// https://leetcode.com/problems/find-all-duplicates-in-an-array/

package main

import (
	"fmt"
)

func findDuplicates(nums []int) []int {
	r := make([]int, 0)
	l := len(nums)
	i := 0
	for i < l {
		if nums[i] > 0 {
			a := nums[i] - 1
			nums[i] = -1
			for a >= 0 && a < l {
				if nums[a] == 0 {
					r = append(r, a+1)
				}
				j := nums[a]
				nums[a] = 0
				a = j - 1
			}
		}
		i++
	}
	return r
}

func main() {
	fmt.Println(findDuplicates([]int{1, 2, 3, 4, 4, 5, 5, 9, 9}))
	fmt.Println(findDuplicates([]int{1,3,3}))
}
