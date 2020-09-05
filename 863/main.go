package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNode1 struct {
	Val    int
	Parent *TreeNode1
	Left   *TreeNode1
	Right  *TreeNode1
}

func NewTree(node *TreeNode, parent *TreeNode1, target *TreeNode, newTarget *TreeNode1) *TreeNode1 {
	if node == nil {
		return nil
	}
	newNode := TreeNode1{Val: node.Val, Parent: parent}
	newNode.Left = NewTree(node.Left, &newNode, target, newTarget)
	newNode.Right = NewTree(node.Right, &newNode, target, newTarget)
	if newNode.Val == target.Val {
		*newTarget = newNode
	}
	return &newNode
}

func FindNode(node *TreeNode1, i, k int, l, r, p bool, output *[]int) {
	if node == nil {
		return
	}
	if i > k {
		return
	}
	if i == k {
		*output = append(*output, node.Val)
	}
	if l {
		FindNode(node.Left, i+1, k, true, true, false, output)
	}
	if r {
		FindNode(node.Right, i+1, k, true, true, false, output)
	}
	if p {
		if node.Parent != nil {
			if node.Parent.Left != nil && node.Parent.Left.Val == node.Val {
				FindNode(node.Parent, i+1, k, false, true, p, output)
			} else if node.Parent.Right != nil && node.Parent.Right.Val == node.Val {
				FindNode(node.Parent, i+1, k, true, false, p, output)
			} else {
			}
		}

	}

}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	var newTarget TreeNode1
	NewTree(root, nil, target, &newTarget)
	output := make([]int, 0)
	if K == 0 {
		output = append(output, newTarget.Val)
	}
	FindNode(newTarget.Left, 1, K, true, true, false, &output)
	FindNode(newTarget.Right, 1, K, true, true, false, &output)
	if newTarget.Parent != nil {
		if newTarget.Parent.Left != nil && newTarget.Parent.Left.Val == newTarget.Val {
			FindNode(newTarget.Parent, 1, K, false, true, true, &output)
		} else if newTarget.Parent.Right != nil && newTarget.Parent.Right.Val == newTarget.Val {
			FindNode(newTarget.Parent, 1, K, true, false, true, &output)
		}
	}
	return output
}

func main() {
	target := &TreeNode{Val: 5, Left: &TreeNode{Val: 6, Right: &TreeNode{Val: 18}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}
	tree := &TreeNode{Val: 3, Left: target, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}}
	root := &TreeNode{Val: 9, Left: &TreeNode{Val: 11, Right: &TreeNode{Val: 19, Left: tree}}}
	// root := tree
	distances := distanceK(root, target, 2)
	fmt.Println(distances)

}
