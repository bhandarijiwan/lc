// https://leetcode.com/problems/interval-list-intersections/
package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func checkOverlap(a, b []int) bool {
	if b[0] > a[1] || b[1] < a[0] {
		return false
	}
	return true
}

func overlap(a, b []int) []int {
	return []int{max(a[0], b[0]), min(a[1], b[1])}
}

func intervalIntersection(A [][]int, B [][]int) [][]int {
	la, lb := len(A), len(B)
	r := make([][]int, 0)
	for i, j := 0, 0; i < la && j < lb; {
		intervalA, intervalB := A[i], B[j]
		if checkOverlap(intervalA, intervalB) {
			r = append(r, overlap(intervalA, intervalB))
		}
		if intervalA[1] < intervalB[1] {
			i = i + 1
		} else if intervalB[1] < intervalA[1] {
			j = j + 1
		} else {
			i, j = i+1, j+1
		}
	}
	return r
}

func main() {
	A := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	B := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	fmt.Println(intervalIntersection(A, B))
}
