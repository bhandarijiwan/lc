// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode, level int, output *[][]int) {
	if root == nil {
		return
	}
	if level < len(*output) {
		a := *output
		a[level] = append(a[level], root.Val)
	} else {
		*output = append(*output, []int{root.Val})
	}
	levelOrder(root.Left, level+1, output)
	levelOrder(root.Right, level+1, output)
}

func levelOrderBottom1(root *TreeNode) [][]int {
	output := make([][]int, 0)
	levelOrder(root, 0, &output)
	i, j := 0, len(output)-1
	for j > i {
		output[i], output[j] = output[j], output[i]
		j = j - 1
		i = i + 1
	}
	return output
}

func levelOrderBottom(root *TreeNode) [][]int {
	output := make([][]int, 0)
	q := make([]*TreeNode, 0)
	if root != nil {
		q = append(q, root)
	}
	for len(q) > 0 {
		l := len(q)
		a := make([]int, 0)
		for i := 0; i < l; i++ {
			a = append(a, q[i].Val)
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		output = append(output, a)
		q = q[l:]
	}
	i, j := 0, len(output)-1
	for j > i {
		output[i], output[j] = output[j], output[i]
		j = j - 1
		i = i + 1
	}
	return output
}

func main() {
	// t := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(levelOrderBottom(nil))
}
