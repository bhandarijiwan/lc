// https://leetcode.com/problems/unique-paths/

package main

import (
	"fmt"
)

func solve(m, n, i, j int, r [][]int) int {
	if m == i && n == j {
		return 1
	}
	if r[i][j] > 0 {
		return r[i][j]
	}
	right := 0
	if i+1 <= m && j <= n {
		right = solve(m, n, i+1, j, r)
	}
	down := 0
	if i <= m && j+1 <= n {
		down = solve(m, n, i, j+1, r)
	}
	r[i][j] = right + down
	return r[i][j]
}

func uniquePaths(m int, n int) int {
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		r[i] = make([]int, m)
	}
	return solve(n-1, m-1, 0, 0, r)
}

func main() {
	fmt.Println(uniquePaths(7, 3))
}
