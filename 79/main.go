// https://leetcode.com/problems/word-search/
/**/
package main

import (
	"fmt"
)

func index(a byte) byte {
	if a >= 'A' && a <= 'Z' {
		return a - 'A'
	}
	return a - 'a'
}

func backTrack(board [][]byte, visited [][]bool, word string, i, j, k, m, n, p int) bool {
	if word[k] != board[i][j] {
		return false
	}
	visited[i][j] = true
	v := false
	if k == p-1 {
		v = true
	}
	if j+1 < n && !visited[i][j+1] && k+1 < p {
		v = backTrack(board, visited, word, i, j+1, k+1, m, n, p)
	}
	if !v && j > 0 && !visited[i][j-1] && k+1 < p {
		v = backTrack(board, visited, word, i, j-1, k+1, m, n, p)
	}
	if !v && i+1 < m && !visited[i+1][j] && k+1 < p {
		v = backTrack(board, visited, word, i+1, j, k+1, m, n, p)
	}
	if !v && i > 0 && !visited[i-1][j] && k+1 < p {
		v = backTrack(board, visited, word, i-1, j, k+1, m, n, p)
	}
	visited[i][j] = false
	return v
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}
	v := false
	m, n, p := len(board), len(board[0]), len(word)
	if p < 1 {
		return false
	}
outer:
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v = backTrack(board, visited, word, i, j, 0, m, n, p)
			if v {
				break outer
			}
		}
	}
	return v
}

func main() {
	input := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(input, "ABCCED"))
	fmt.Println(exist(input, "ABCB"))
	fmt.Println(exist(input, "SEE"))
}
