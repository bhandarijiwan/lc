// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func buildTreeHelper(preOrder []int, inOrderMap map[int]int, startIndex int) *TreeNode {
	if len(preOrder) < 1 {
		return nil
	}
	val := preOrder[0]
	valIndex, _ := inOrderMap[val]
	count := valIndex - startIndex
	node := &TreeNode{Val: val}
	leftLimit := min(1+count, len(preOrder))
	node.Left = buildTreeHelper(preOrder[1:leftLimit], inOrderMap, startIndex)
	node.Right = buildTreeHelper(preOrder[leftLimit:], inOrderMap, startIndex+leftLimit)
	return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	inOrderMap := make(map[int]int)
	for index, num := range inorder {
		inOrderMap[num] = index
	}
	return buildTreeHelper(preorder, inOrderMap, 0)
}

func main() {
	// preOrder := []int{1, 2, 3, 4}
	// inOrder := []int{4, 3, 2, 1}
	// root := buildTree(preOrder, inOrder)
	// fmt.Println(root.Val)
	// fmt.Println(root.Left.Val)
	// fmt.Println(root.Left.Left.Val)
	// fmt.Println(root.Left.Left.Left.Val)

	// preOrder := []int{3, 9, 20, 15, 7}
	// inOrder := []int{9, 3, 15, 20, 7}
	// root := buildTree(preOrder, inOrder)
	// fmt.Println(root.Val)
	// fmt.Println(root.Left.Val, root.Right.Val)
	// fmt.Println(root.Right.Left.Val, root.Right.Right.Val)

	// preOrder := []int{1, 2, 3, 4}
	// inOrder := []int{2, 3, 1, 4}
	// root := buildTree(preOrder, inOrder)
	// fmt.Println(root.Val)
	// fmt.Println(root.Left.Val, root.Right.Val)

	// preOrder := []int{1, 2, 3, 4}
	// inOrder := []int{2, 4, 3, 1}
	// root := buildTree(preOrder, inOrder)
	// fmt.Println(root.Val)
	// fmt.Println(root.Left.Val)
	// fmt.Println(root.Left.Right.Val)

	preOrder := []int{1, 2, 3, 4, 5, 6}
	inOrder := []int{1, 3, 2, 5, 4, 6}
	root := buildTree(preOrder, inOrder)
	fmt.Println(root.Val)
	fmt.Println(root.Right.Val)
	fmt.Println(root.Right.Left.Val)
	fmt.Println(root.Right.Right.Val)
	fmt.Println(root.Right.Right.Left.Val)
	fmt.Println(root.Right.Right.Right.Val)

}
