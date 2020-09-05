// https://leetcode.com/problems/rotate-image/

package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	m := len(matrix)
	k := (m - 1) / 2
	for i := 0; i <= k; i++ {
		h := m - 1 - i
		for j := i; j < m-1-i; j++ {
			r, c := j, h
			a1 := matrix[r][c]
			matrix[r][c] = matrix[i][j]
			r, c = c, m-j-1
			a2 := matrix[r][c]
			matrix[r][c] = a1
			r, c = m-j-1, i
			a3 := matrix[r][c]
			matrix[r][c] = a2
			matrix[i][j] = a3
		}
	}
}

func main() {
	input := [][]int{{1, 2}, {3, 4}}
	rotate(input)
	fmt.Println(input)
	input = [][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	rotate(input)
	fmt.Println(input)
	input = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(input)
	input = [][]int{
		{15, 13, 18, 19, 20},
		{1, 4, 5, 6, 7},
		{8, 9, 10, 11, 12},
		{13, 14, 15, 16, 17},
		{18, 19, 20, 21, 22},
	}
	rotate(input)
	fmt.Println(input)
}
