// https://leetcode.com/problems/edit-distance/
// #Levenshtein distance
package main

import (
	"fmt"
)

func min(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		return c
	}
	return m
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m < 1 {
		return n
	}
	if n < 1 {
		return m
	}
	x := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		x[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		x[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		x[0][i] = i
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			c := 1
			if word1[i-1] == word2[j-1] {
				c = 0
			}
			x[i][j] = min(x[i-1][j]+1, x[i][j-1]+1, x[i-1][j-1]+c)
		}
	}
	return x[m][n]
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}
