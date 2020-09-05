//https://leetcode.com/problems/single-number/
package main

import "fmt"

func singleNumber1(nums []int) int {
	m := make(map[int]bool)
	s := 0
	for _, c := range nums {
		if _, ok := m[c]; ok {
			s = s - c
		} else {
			s = s + c
			m[c] = true
		}
	}
	return s
}

func singleNumber(nums []int) int {
	s := 0
	for _, n := range nums {
		s = s^n
	}
	return s
}

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}
