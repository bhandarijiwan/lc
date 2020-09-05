//https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(t *TreeNode) []int {
	if t == nil {
		return []int{}
	}
	m := []int{t.Val}
	l := printTree(t.Left)
	r := printTree(t.Right)
	for _, a := range l {
		m = append(m, a)
	}
	for _, a := range r {
		m = append(m, a)
	}
	fmt.Println("Root =", t.Val)
	fmt.Println("Left =", l)
	fmt.Println("Right =", r)

	return m
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("%v", printTree(t))
}

func createBST(nums []int, l, r int) *TreeNode {
	if r < l {
		return nil
	}
	m := l + (r-l)/2
	n := &TreeNode{Val: nums[m]}
	n.Left = createBST(nums, l, m-1)
	n.Right = createBST(nums, m+1, r)
	return n
}

func sortedArrayToBST(nums []int) *TreeNode {
	return createBST(nums, 0, len(nums)-1)
}
func main() {
	fmt.Println(sortedArrayToBST([]int{-10,-2,1}))
}
