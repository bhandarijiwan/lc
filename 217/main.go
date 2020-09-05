//https://leetcode.com/problems/contains-duplicate/
package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}
		m[n]++
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1,2,3,4}))
}
