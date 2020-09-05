// https://leetcode.com/problems/house-robber-iii/
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

func helper(root *TreeNode, parentRobbed bool) int {
	if root == nil {
		return 0
	}
	skipCurrent :=  helper(root.Left, false) + helper(root.Right, false)
	if parentRobbed {
		return skipCurrent
	}
	robCurrent := root.Val + helper(root.Left, true) + helper(root.Right, true)
	if skipCurrent > robCurrent {
		return skipCurrent
	}
	return robCurrent

}

func rob(root *TreeNode) int {
	r1 := helper(root, true)
	r2 := helper(root, false)
	if r2 > r1 {
		return r2
	}
	return r1
}

func main() {
	tree := &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}}
	fmt.Println(rob(tree))
	tree = &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 1}}}
	fmt.Println(rob(tree))
	fmt.Println(rob(nil))
	tree = &TreeNode{Val: 4, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}}
	fmt.Println(rob(tree))
}
