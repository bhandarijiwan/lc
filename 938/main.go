//https://leetcode.com/problems/range-sum-of-bst/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Add(item int) {
	n := t
	d := -1
	for true {
		if item > n.Val {
			if n.Right != nil {
				n = n.Right
			} else {
				d = 1
				break
			}
		} else {
			if n.Left != nil {
				n = n.Left
			} else {
				d = 0
				break
			}
		}
	}

	if d >= 0 {
		if d > 0 {
			n.Right = &TreeNode{Val: item}
		} else {
			n.Left = &TreeNode{Val: item}
		}
	} else {
		n.Val = item
	}

}

func (t *TreeNode) String() string {
	return ""
}

func overlaps(n ...int) bool {
	l := n[0]
	r := n[1]
	for _, i := range n[2:] {
		if i >= l && i <= r {
			return true
		}
	}
	if n[2] < l && n[3] > r {
		return true
	}
	return false
}

func rangeSumBST(root *TreeNode, l int, R int) int {
	if root == nil {
		return 0
	}
	v := 0
	if root.Val >= l && root.Val <= R {
		v = root.Val
	}
	left := 0
	if root.Left != nil {
		if overlaps(l, R, root.Val, root.Left.Val) {
			left = rangeSumBST(root.Left, l, R)
		}
	}
	right := 0
	if root.Right != nil {
		if overlaps(l, R, root.Val, root.Right.Val) {
			right = rangeSumBST(root.Right, l, R)
		}
	}
	return v + left + right
}

func newTree(items []int) *TreeNode {
	tree := &TreeNode{}
	for _, i := range items {
		if i >= 0 {
			tree.Add(i)
		}
	}
	return tree
}

func main() {

	tree := newTree([]int{10, 5, 15, 3, 7, -1, 18})
	fmt.Println(rangeSumBST(tree, 7, 15))
	tree2 := newTree([]int{10, 5, 15, 3, 7, 13, 18, 1, -1, 6})
	fmt.Println(rangeSumBST(tree2, 6, 10))
	tree3 := newTree([]int{18, 9, 27, 6, 15, 24, 30, 3, -1, 12, -1, 21})
	fmt.Println(rangeSumBST(tree3, 18, 24))
	fmt.Println(overlaps(7, 15, 5, 20))
}
