// https://leetcode.com/problems/kth-smallest-element-in-a-bst/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack struct {
	data []*TreeNode
	t    int
}

func NewStack() *stack {
	return &stack{data: make([]*TreeNode, 2), t: -1}
}

func (s *stack) Push(node *TreeNode) {
	s.t = s.t + 1
	l := len(s.data)
	if l <= s.t {
		newData := make([]*TreeNode, 2*l)
		copy(newData, s.data)
		s.data = newData
	}
	s.data[s.t] = node
}

func (s *stack) Pop() *TreeNode {
	if s.t >= 0 {
		a := s.data[s.t]
		s.t = s.t - 1
		return a
	}
	return nil
}

func (s *stack) isEmpty() bool {
	return s.t < 0
}
func (s *stack) Peek() *TreeNode {
	if s.t >= 0 {
		return s.data[s.t]
	}
	return nil
}

func kthSmallest1(root *TreeNode, k int) int {
	s := NewStack()
	t := root
	for t != nil {
		for t != nil {
			s.Push(t)
			t = t.Left
		}
		for !s.isEmpty() {
			k = k - 1
			t = s.Pop()
			if k == 0 {
				return t.Val
			}
			t = t.Right
			if t != nil {
				break
			}
		}
	}
	return 0
}

func main() {
	t := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}
	fmt.Println(kthSmallest1(t, 1))
	t1 := &TreeNode{Val: 5,
		Left: &TreeNode{Val: 3,
			Right: &TreeNode{Val: 4},
			Left: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 1}}},
		Right: &TreeNode{Val: 6}}
	fmt.Println(kthSmallest1(t1, 3))
	fmt.Println(kthSmallest1(&TreeNode{Val: 3}, 3))
}
