// https://leetcode.com/problems/minimum-path-sum/

package main

import (
	"fmt"
)

// go:inline
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// go:inline
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	path := make([][]int, len(grid))
	for i := 0; i < m; i++ {
		path[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			d := grid[i][j]
			if i+1 < m {
				d += path[i+1][j]
			}
			if j+1 < n {
				if i+1 < m {
					d = min(d, grid[i][j]+path[i][j+1])
				} else {
					d += path[i][j+1]
				}
			}
			path[i][j] = d
		}
	}
	return path[0][0]
}

func main() {
	// input := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	input := [][]int{{9, 9, 0, 8, 9, 0, 5, 7, 2, 2, 7, 0, 8, 0, 2, 4, 8}, {4, 4, 2, 7, 6, 0, 9, 7, 3, 2, 5, 4, 6, 5, 4, 8, 7}, {4, 9, 7, 0, 7, 9, 2, 4, 0, 2, 4, 4, 6, 2, 8, 0, 7}, {7, 7, 9, 6, 6, 4, 8, 4, 8, 7, 9, 4, 7, 6, 9, 6, 5}, {1, 3, 7, 5, 7, 9, 7, 3, 3, 3, 8, 3, 6, 5, 0, 3, 6}, {7, 1, 0, 7, 5, 0, 6, 6, 5, 3, 2, 6, 0, 0, 9, 5, 7}, {6, 5, 6, 3, 8, 1, 8, 6, 4, 4, 3, 4, 9, 9, 3, 3, 1}, {1, 0, 2, 9, 7, 9, 3, 1, 7, 5, 1, 8, 2, 8, 4, 7, 6}, {9, 6, 7, 7, 4, 1, 4, 0, 6, 5, 1, 9, 0, 3, 2, 1, 7}, {2, 0, 8, 7, 1, 7, 4, 3, 5, 6, 1, 9, 4, 0, 0, 2, 7}, {9, 8, 1, 3, 8, 7, 1, 2, 8, 3, 7, 3, 4, 6, 7, 6, 6}, {4, 8, 3, 8, 1, 0, 4, 4, 1, 0, 4, 1, 4, 4, 0, 3, 5}, {6, 3, 4, 7, 5, 4, 2, 2, 7, 9, 8, 4, 5, 6, 0, 3, 9}, {0, 4, 9, 7, 1, 0, 7, 7, 3, 2, 1, 4, 7, 6, 0, 0, 0}}
	fmt.Println(minPathSum(input))
}
