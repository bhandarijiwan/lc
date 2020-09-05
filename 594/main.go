//https://leetcode.com/problems/longest-harmonious-subsequence/

package main

import "fmt"

func findLHS(nums []int) int {
	m := make(map[int]int)
	l := len(nums)
	p := 0
	for i := 0; i < l; i++ {
		m[nums[i]]++
		if a, ok := m[nums[i]+1]; ok {
			b := a + m[nums[i]]
			if b > p {
				p = b
			}
		}
		if a, ok := m[nums[i]-1]; ok {
			b := a + m[nums[i]]
			if b > p {
				p = b
			}
		}
	}
	return p
}
func main() {
	fmt.Println(findLHS([]int{1,3,2,2,5,2,3,7}))
}
