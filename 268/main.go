//https://leetcode.com/problems/missing-number/

package main

import "fmt"

func missingNumber(nums []int) int {
	s := 0
	l := len(nums)
	es := (l * (l + 1)) / 2
	for _, n := range nums {
		s += n
	}
	return es - s
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
