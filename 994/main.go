package main

import (
	"fmt"
)

type Grid [][]int

func CloneEmpty(g Grid) Grid {
	m := len(g)
	n := len(g[0])
	newGrid := make([][]int, m)
	for i := 0; i < m; i++ {
		newGrid[i] = make([]int, n)
	}
	return newGrid
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func process(grid Grid, i, j, m, n, d int, ng Grid) {
	v := grid[i][j]
	grid[i][j] = 2
	if ng[i][j] > 0 {
		ng[i][j] = min(ng[i][j], d)
	} else {
		ng[i][j] = d
	}
	if i > 0 && grid[i-1][j] == 1 {
		process(grid, i-1, j, m, n, d+1, ng)
	}
	if i < m-1 && grid[i+1][j] == 1 {
		process(grid, i+1, j, m, n, d+1, ng)
	}
	if j > 0 && grid[i][j-1] == 1 {
		process(grid, i, j-1, m, n, d+1, ng)
	}
	if j < n-1 && grid[i][j+1] == 1 {
		process(grid, i, j+1, m, n, d+1, ng)
	}
	grid[i][j] = v
}

func orangesRotting(grid [][]int) int {
	if len(grid) < 1 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	ng := CloneEmpty(grid)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				process(grid, i, j, m, n, 0, ng)
			}
		}
	}
	maxCount := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && ng[i][j] < 1 {
				return -1
			} else {
				if ng[i][j] > maxCount {
					maxCount = ng[i][j]
				}
			}
		}
	}
	return maxCount
}

func main() {
	input := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Println(orangesRotting(input))
	input = [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	fmt.Println(orangesRotting(input))
	input = [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 2}}
	fmt.Println(orangesRotting(input))
	input = [][]int{{0, 2}}
	fmt.Println(orangesRotting(input))
	input = [][]int{{}}
	fmt.Println(orangesRotting(input))
	input = [][]int{{0}}
	fmt.Println(orangesRotting(input))
	input = [][]int{{2}}
	fmt.Println(orangesRotting(input))
	input = [][]int{{1, 2}}
	fmt.Println(orangesRotting(input))

	input = [][]int{{1}, {2}, {1}, {1}}
	fmt.Println(orangesRotting(input))

	input = [][]int{{0, 0, 1, 2}, {2, 0, 1, 1}}
	fmt.Println(orangesRotting(input))

}
