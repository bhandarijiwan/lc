// https://leetcode.com/problems/find-largest-value-in-each-tree-row/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLargest(root *TreeNode, level int, output *[]int) {
	if root == nil {
		return
	}
	l := len(*output)
	if level < l {
		if root.Val > (*output)[level] {
			(*output)[level] = root.Val
		}
	} else {
		*output = append(*output, root.Val)
	}
	findLargest(root.Left, level+1, output)
	findLargest(root.Right, level+1, output)
}

func largestValues(root *TreeNode) []int {
	output := make([]int, 0)
	findLargest(root, 0, &output)
	return output
}

func main() {
	t := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2,
			Right: &TreeNode{Val: 9}}}
	fmt.Println(largestValues(t))
}
