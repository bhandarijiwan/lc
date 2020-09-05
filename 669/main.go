// https://leetcode.com/problems/trim-a-binary-search-tree/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	if root.Val < L || root.Val > R {
		if root.Val > R {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

func main() {
	t := &TreeNode{Val: 3, Right: &TreeNode{Val: 4}, Left: &TreeNode{Val: 0, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}}
	ot := trimBST(t, 1, 3)
	fmt.Println(ot.Val, ot.Left, ot.Left.Left)

	tt := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}, Left: &TreeNode{Val: 0}}
	ot = trimBST(tt, 3, 4)

}
