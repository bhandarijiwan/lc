// https://leetcode.com/problems/validate-binary-search-tree/
package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	output := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		output = append(output, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	for i := 1; i < len(output); i++ {
		if output[i] <= output[i-1] {
			return false
		}
	}
	return true
}

func main() {
	tree := &TreeNode{Val: 10, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 20}}}
	fmt.Println(isValidBST(tree))
	tree = &TreeNode{Val: 10, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 11}, Right: &TreeNode{Val: 20}}}
	fmt.Println(isValidBST(tree))
	tree = &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	fmt.Println(isValidBST(tree))
	tree = &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	fmt.Println(isValidBST(tree))
	tree = &TreeNode{Val: 3, Right: &TreeNode{Val: 30, Left: &TreeNode{Val: 10, Right: &TreeNode{Val: 15, Right: &TreeNode{Val: 45}}}}}
	fmt.Println(isValidBST(tree))
}
