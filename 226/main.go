// https://leetcode.com/problems/invert-binary-tree/
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invert(t *TreeNode) {
	if t == nil {
		return
	}
	t.Left, t.Right = t.Right, t.Left
	invert(t.Left)
	invert(t.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		invert(root)
	}
	return root
}

func main() {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}
	invertTree(tree)
	fmt.Println(tree.Val)
	fmt.Println(tree.Left.Val, ",", tree.Right.Val)
	fmt.Println(tree.Left.Left.Val, ",", tree.Left.Right.Val, ",", tree.Right.Left.Val, ",", tree.Right.Right.Val)
	tree = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	invertTree(tree)
	fmt.Println()
	fmt.Println(tree.Val)
	fmt.Println(tree.Left, tree.Right.Val)
	tree = &TreeNode{Val: 9, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 8}}}
	invertTree(tree)
	fmt.Println()
	fmt.Println(tree.Val)
	fmt.Println(tree.Left.Val)
	fmt.Println(tree.Left.Right.Val)
}
