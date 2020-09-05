// https://leetcode.com/problems/unique-binary-search-trees/
package main

import (
	"fmt"
)

func makeTree(r, l, h, p, c int) int {
	if l+1 >= r && r+1 >= h {
		return 1
	}
	left := 0
	for i := l + 1; i < r; i++ {
		left += makeTree(i, l, r-1, r, c)
	}
	right := 0
	for i := r + 1; i <= h; i++ {
		right += makeTree(i, r, h, r, c)
	}
	if left > 0 && right > 0 {
		return left * right
	}
	if right > left {
		return right
	}
	return left
}

func numTrees(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += makeTree(i, 0, n, i, n-1)
	}
	return sum
}

func main() {
	fmt.Println(numTrees(0))
	fmt.Println(numTrees(6))
	fmt.Println(numTrees(5))
	fmt.Println(numTrees(4))
	fmt.Println(numTrees(3))
}
