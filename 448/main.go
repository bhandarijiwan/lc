//https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] != 0 {
			a := nums[i]
			for a != 0 {
				b := nums[a-1]
				nums[a-1] = 0
				a = b
			}
		}
	}
	m := make([]int, 0)
	for i, c := range nums {
		if c != 0 {
			m = append(m, i+1)
		}
	}
	return m
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
