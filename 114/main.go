// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
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

func helper(root *TreeNode, prev **TreeNode) {
	if root == nil {
		return
	}
	(**prev).Right = root
	(**prev).Left = nil
	*prev = root
	right := root.Right
	helper(root.Left, prev)
	helper(right, prev)
}

func flatten(root *TreeNode) {
	t := &TreeNode{}
	helper(root, &t)
}

func main() {
	tree := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 5,
			Right: &TreeNode{Val: 6},
		},
	}
	flatten(tree)
	k := tree
	for k != nil {
		fmt.Println(k.Val)
		k = k.Right
	}
	tree = &TreeNode{Val: 9, Right: &TreeNode{Val: 19}, Left: &TreeNode{Val: 89}}
	flatten(tree)
	k = tree
	for k != nil {
		fmt.Println(k.Val)
		k = k.Right
	}

}
