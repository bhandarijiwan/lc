// https://leetcode.com/problems/maximum-depth-of-binary-tree/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxdepth(root *TreeNode, d int, output *int) {
	if root == nil {
		return
	}
	if d > *output {
		*output = d
	}
	maxdepth(root.Left, d+1, output)
	maxdepth(root.Right, d+1, output)
}

func maxDepth(root *TreeNode) int {
	d, output := 1, 0
	maxdepth(root, d, &output)
	return output
}

func main() {
	tree := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 8}}}}
	fmt.Println(maxDepth(tree))
}
