// https://leetcode.com/problems/maximal-rectangle/

package main

import (
	"fmt"
	"strings"
)

func PrintByteMatrix(m [][]byte) {
	var sb strings.Builder
	for i := 0; i < len(m); i++ {
		sb.WriteString(string(m[i]))
		sb.WriteByte('\n')
	}
	fmt.Println(sb.String())
}

func PrintIntMatrix(m [][]int) {
	var sb strings.Builder
	for i := 0; i < len(m); i++ {
		sb.WriteString(fmt.Sprintln(m[i]))
	}
	fmt.Println(sb.String())
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m < 1 {
		return 0
	}
	n := len(matrix[0])
	l, u := make([][]int, m), make([][]int, m)
	for i := 0; i < m; i++ {
		l[i] = make([]int, n)
		u[i] = make([]int, n)
	}
	mxy := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if i > 0 && matrix[i-1][j] == '1' {
					u[i][j] += u[i-1][j] + 1
				}
				if j > 0 && matrix[i][j-1] == '1' {
					l[i][j] += l[i][j-1] + 1
				}
				x, y := l[i][j]+1, u[i][j]+1
				if x > y {
					mx := x
					p := i - 1
					for p >= 0 && matrix[p][j] == '1' {
						mx = max(min(l[p][j]+1, x)*((i-p)+1), mx)
						x = min(l[p][j]+1, x)
						p--
					}
					mxy = max(mxy, mx)
					// fmt.Println("max = ", mxy, "i, j = ", i, j)
				} else {
					my := y
					p := j - 1
					for p >= 0 && matrix[i][p] == '1' {
						my = max(min(u[i][p]+1, y)*((j-p)+1), my)
						y = min(u[i][p]+1, y)
						p--
					}
					mxy = max(mxy, my)
				}
			}
		}
	}
	// PrintByteMatrix(matrix)
	// fmt.Println("Left Matrix")
	// PrintIntMatrix(l)
	// fmt.Println("Up Matrix")
	// PrintIntMatrix(u)
	// fmt.Println(mxy)
	return mxy
}

func main() {
	input := [][]byte{{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}

	fmt.Println(maximalRectangle(input))

	input = [][]byte{{'1', '1', '1'}}
	fmt.Println(maximalRectangle(input))

	input = [][]byte{}
	fmt.Println(maximalRectangle(input))

	input = [][]byte{
		{'1', '0', '1', '0', '1'},
		{'1', '0', '0', '1', '1'},
		{'1', '1', '0', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	fmt.Println(maximalRectangle(input))

	input = [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '0', '1', '1'},
		{'1', '1', '0', '1', '1'},
		{'1', '0', '0', '1', '1'}}
	fmt.Println(maximalRectangle(input))
}
