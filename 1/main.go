//https://leetcode.com/problems/two-sum/

package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		c := target - n
		if _, ok := m[c]; ok {
			return []int{m[c], i}
		}
		m[n] = i
	}
	panic("not found")
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
