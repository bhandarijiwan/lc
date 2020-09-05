//https://leetcode.com/problems/squares-of-a-sorted-array/

package main

import "fmt"

func mergeTwoSortedArrays(a, b []int) []int {
	c := make([]int, len(a)+len(b))
	ai := len(a) - 1
	bi := 0
	i := 0
	for ai >= 0 && bi < len(b) {
		if a[ai] < b[bi] {
			c[i] = a[ai]
			ai--
		} else {
			c[i] = b[bi]
			bi++
		}
		i++
	}
	for ai >= 0 {
		c[i] = a[ai]
		ai--
		i++
	}
	for bi < len(b) {
		c[i] = b[bi]
		bi++
		i++
	}
	return c
}
func sortedSquares(A []int) []int {
	p := -1
	for i := 0; i < len(A); i++ {
		if A[i] >= 0 && p < 0 {
			p = i
		}
		A[i] = A[i] * A[i]
	}
	if p >= 0 {
		return mergeTwoSortedArrays(A[0:p], A[p:len(A)])
	}
	return A
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
