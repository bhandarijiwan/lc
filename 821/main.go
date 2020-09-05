//https://leetcode.com/problems/shortest-distance-to-a-character/
package main

import (
	"fmt"
)

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -1 * n
}

func find(nums []int, n int) int {
	l := 0
	r := len(nums) - 1
	var m int
	for r >= l {
		m = l + (r-l)/2
		if n >= nums[m] {
			if n > nums[m] {
				l = m + 1
			} else {
				return m
			}
		} else {
			r = m - 1
		}
	}
	if l >= len(nums) {
		return r
	}
	if r < 0 {
		return l
	}
	if abs(nums[l]-n) >= abs(nums[r]-n) {
		return r
	}
	return l
}

func shortestToChar1(S string, C byte) []int {
	l := len(S)
	d := make([]int, l)
	indices := make([]int, 0)
	ci := int32(C)
	for i, c := range S {
		if c == ci {
			indices = append(indices, i)
		}
	}
	for i := 0; i < l; i++ {
		a := find(indices, i)
		d[i] = abs(i - indices[a])
	}
	return d
}

func shortestToChar(S string, C byte) []int {
	d := make([]int, len(S))
	ci := int32(C)
	p := -len(S)
	for i, c := range S {
		if c == ci {
			p = i
		}
		d[i] = i - p
	}
	for i := p; i >= 0; i-- {
		if int32(S[i]) == ci {
			p = i
		}
		if d[i] > p-i {
			d[i] = p - i
		}
	}
	return d
}

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}
