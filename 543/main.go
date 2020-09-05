//https://leetcode.com/problems/diameter-of-binary-tree/
package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterSubTree(t *TreeNode, a []int) int {
	if t == nil {
		return 0
	}
	l := diameterSubTree(t.Left, a)
	r := diameterSubTree(t.Right, a)
	if l+r > a[0] {
		a[0] = l + r
	}
	if l > r {
		return 1 + l
	}
	return 1 + r
}

func diameterOfBinaryTree(root *TreeNode) int {
	a := make([]int, 1)
	diameterSubTree(root, a)
	return a[0]
}

func main() {

	t := &TreeNode{
		Val: 4,
		Right: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -9,
				Left: &TreeNode{
					Val:   6,
					Left:  &TreeNode{Val: 0, Right: &TreeNode{Val: -1}},
					Right: &TreeNode{Val: 6, Left: &TreeNode{Val: -4}},
				},
				Right: &TreeNode{
					Val:   -7,
					Left:  &TreeNode{Val: -6, Left: &TreeNode{Val: 5}},
					Right: &TreeNode{Val: -6, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: -2}}},
				},
			},
			Right: &TreeNode{Val: -3, Left: &TreeNode{Val: -4}},
		},
		Left: &TreeNode{
			Val: -7,
		},
	}
	fmt.Println(diameterOfBinaryTree(t))
}
