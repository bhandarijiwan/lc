// https://leetcode.com/problems/binary-tree-level-order-traversal/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lb(root *TreeNode, level int, output *[][]int) {
	if root == nil {
		return
	}
	if level < len(*output) {
		a := *output
		a[level] = append(a[level], root.Val)
	} else {
		*output = append(*output, []int{root.Val})
	}
	lb(root.Left, level+1, output)
	lb(root.Right, level+1, output)
}

func levelOrder(root *TreeNode) [][]int {
	output := make([][]int, 0)
	lb(root, 0, &output)
	return output
}

func main() {
	t := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(levelOrder(t))
}
