// https://leetcode.com/problems/maximal-square/

package main

import (
	"fmt"
)

func PrintWindow(window [][]byte, r, c, m, n int) {
	var s string
	for i := r; i < m; i++ {
		for j := c; j < n; j++ {
			s += string(window[i][j])
		}
		s += "\n"
	}
	fmt.Println(s)
}

func CheckWindow(window [][]byte, r, c, m, n int) bool {
	for i := r; i < m; i++ {
		for j := c; j < n; j++ {
			if window[i][j] != '1' {
				return false
			}
		}
	}
	return true
}

func maximalSquare1(matrix [][]byte) int {
	m := len(matrix)
	if m < 1 {
		return 0
	}
	n := len(matrix[0])
	d := m
	if n < d {
		d = n
	}
	l := 0
nextLevel:
	for l < d {
		for i := 0; i < m-l; i += 1 {
			for j := 0; j < n-l; j += 1 {
				if CheckWindow(matrix, i, j, i+l+1, j+l+1) {
					l++
					continue nextLevel
				}
			}
		}
		break
	}
	return l * l
}

func min(a, b, c int) int {
	if b < a {
		a = b
	}
	if c < a {
		a = c
	}
	return a
}

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m < 1 {
		return 0
	}
	n := len(matrix[0])
	r := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		r[i] = make([]int, len(matrix[i]))
	}
	maxSize := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				r[i][j] = 1
				if i-1 >= 0 && j-1 >= 0 {
					if matrix[i][j-1] == '1' && matrix[i-1][j] == '1' {
						r[i][j] = min(r[i-1][j-1]+1, r[i-1][j]+1, r[i][j-1]+1)
					}
				}
				if r[i][j] > maxSize {
					maxSize = r[i][j]
				}
			} else {
				r[i][j] = 0
			}
		}
	}
	// for _, bytes := range matrix {
	// 	fmt.Println(string(bytes))
	// }
	// fmt.Println(r)
	return maxSize * maxSize
}

func main() {
	input := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	fmt.Println(maximalSquare(input))
	input = [][]byte{{'2', '0', '2', '0', '0'}, {'2', '0', '2', '2', '2'}, {'2', '2', '2', '2', '2'}, {'2', '0', '0', '2', '0'}}
	fmt.Println(maximalSquare(input))
	input = [][]byte{{'1', '1', '2', '0', '0'}, {'2', '0', '2', '2', '2'}, {'2', '2', '2', '2', '2'}, {'2', '0', '0', '2', '0'}}
	fmt.Println(maximalSquare(input))
	fmt.Println(maximalSquare([][]byte{}))
	input = [][]byte{{'1', '1'}, {'1', '1'}}
	fmt.Println(maximalSquare(input))

	input = [][]byte{{'0', '0', '0', '1'}, {'1', '1', '0', '1'}, {'1', '1', '1', '1'}, {'0', '1', '1', '1'}, {'0', '1', '1', '1'}}
	fmt.Println(maximalSquare(input))

}
