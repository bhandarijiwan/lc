// https://leetcode.com/problems/merge-two-binary-trees/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees1(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// merge the two trees into t1
	if t1 == nil {
		return t2
	}
	if t2 != nil {
		t1.Val += t2.Val
		t1.Left = mergeTrees1(t1.Left, t2.Left)
		t1.Right = mergeTrees1(t1.Right, t2.Right)
	}
	return t1
}

type Stack struct {
	items []*TreeNode
	top   int
}

func NewStack() *Stack {
	return &Stack{items: make([]*TreeNode, 1), top: -1}
}

func (s *Stack) Peek() *TreeNode {
	if s.top >= 0 {
		return s.items[s.top]
	}
	return nil
}
func (s *Stack) Push(item *TreeNode) {
	s.top = s.top + 1
	if s.top >= len(s.items) {
		newItems := make([]*TreeNode, len(s.items)*2)
		copy(newItems, s.items)
		s.items = newItems
	}
	s.items[s.top] = item
}

func (s *Stack) Pop() *TreeNode {
	if s.top >= 0 {
		t := s.items[s.top]
		s.top--
		return t
	}
	return nil
}

func (s *Stack) IsEmpty() bool {
	return s.top < 0
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	s1, s2 := NewStack(), NewStack()
	tree1, tree2 := t1, t2
	for tree1 != nil && tree2 != nil {
		for tree1 != nil && tree2 != nil {
			s1.Push(tree1)
			s2.Push(tree2)
			tree1, tree2 = tree1.Left, tree2.Left
		}
		for !s1.IsEmpty() && !s2.IsEmpty() {
			tr1, tr2 := s1.Pop(), s2.Pop()
			tr1.Val += tr2.Val
			if tr1.Left == nil {
				tr1.Left = tr2.Left
			}
			tree1, tree2 = tr1.Right, tr2.Right
			if tree1 != nil && tree2 != nil {
				break
			}
			if tr2.Right != nil {
				tr1.Right = tr2.Right
			}
		}
	}
	if t1 != nil {
		return t1
	}
	return t2
}

func main() {
	// Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}},
	t1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}
	t2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}}
	mergeTrees(t1, t2)
	fmt.Println(t1.Val)
	fmt.Println(t1.Left.Val)
	fmt.Println(t1.Left.Left.Val)
	// fmt.Println(t1.Left.Left.Val, t1.Left.Right.Val, t1.Right.Left, t1.Right.Right.Val)

}
