// https://leetcode.com/problems/binary-tree-inorder-traversal/

package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	if t == nil {
		return "<nil>"
	}
	return strconv.Itoa(t.Val)
}

func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := inorderTraversal1(root.Left)
	right := inorderTraversal1(root.Right)
	lLeft := len(left)
	lRight := len(right)
	r := make([]int, lLeft+lRight+1)
	copy(r, left)
	r[lLeft] = root.Val
	copy(r[lLeft+1:], right)
	return r
}

func inorderTraversal2(root *TreeNode) []int {
	s := make([]*TreeNode, 1)
	s[0] = root
	i := uint(0) // height count
	for {
		c := uint(1 << i) // nodes count at height i
		leaf := true
		nextLevel := make([]*TreeNode, 1<<(i+1))
		index, j := uint(1<<i-1), 0
		for c > 0 {
			node := s[index]
			if node != nil && (node.Left != nil || node.Right != nil) {
				nextLevel[j] = node.Left
				nextLevel[j+1] = node.Right
				leaf = false
			}
			index = index + 1
			j = j + 2
			c = c - 1
		}
		if leaf {
			break
		}
		s = append(s, nextLevel...)
		i++
	}
	ls := len(s)
	inOrderS := make([]*TreeNode, ls)
	levelCount := i
	i = 0
	for i <= levelCount {
		offset := 1<<(levelCount-i) - 1
		step := 1 << (levelCount - i + 1)
		index, c := uint(1<<i-1), uint(1<<i)
		for c > 0 {
			inOrderS[offset] = s[index]
			offset = offset + step
			index = index + 1
			c = c - 1
		}
		i++
	}
	r := make([]int, 0)
	for i := 0; i < len(inOrderS); i++ {
		if inOrderS[i] != nil {
			r = append(r, inOrderS[i].Val)
		}
	}
	return r
}

type Stack struct {
	data []*TreeNode
	top  int
}

func NewStack() *Stack {
	data := make([]*TreeNode, 2)
	return &Stack{data: data, top: -1}
}

func (s *Stack) Push(item *TreeNode) {
	s.top = s.top + 1
	l := len(s.data)
	if l <= s.top {
		newData := make([]*TreeNode, l*2)
		copy(newData, s.data)
		s.data = newData
	}
	s.data[s.top] = item
}

func (s *Stack) Pop() *TreeNode {
	if s.top >= 0 {
		r := s.data[s.top]
		s.top = s.top - 1
		return r
	}
	panic("Empty stack!!")
}

func (s *Stack) Empty() bool {
	if s.top < 0 {
		return true
	}
	return false
}
func (s *Stack) Peek() *TreeNode {
	if s.top >= 0 {
		return s.data[s.top]
	}
	return nil
}

func inorderTraversal(root *TreeNode) []int {
	s := NewStack()
	t, r := root, make([]int, 0)
	for t != nil {
		for t != nil {
			s.Push(t)
			t = t.Left
		}
		for !s.Empty() {
			n := s.Pop()
			r = append(r, n.Val)
			if n.Right != nil {
				t = n.Right
				break
			}
		}
	}
	return r
}

func main() {
	t := &TreeNode{Val: 2, Right: nil, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 1}}}
	fmt.Println(inorderTraversal(t))
}
