// https://leetcode.com/problems/largest-rectangle-in-histogram/
package main

import (
	"fmt"
)

type s struct {
	i int
	a int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Stack struct {
	d []int
	t int
}

func NewStack() *Stack {
	return &Stack{d: make([]int, 1), t: -1}
}

func (s *Stack) Push(item int) {
	s.t++
	if s.t >= len(s.d) {
		nd := make([]int, len(s.d)*2)
		copy(nd, s.d)
		s.d = nd
	}
	s.d[s.t] = item
}

func (s *Stack) Pop() int {
	if s.t < 0 {
		panic("Stack is empty!!")
	}
	a := s.d[s.t]
	l := len(s.d)
	if s.t <= (l >> 2) {
		s.d = s.d[:(l >> 1)]
	}
	s.t--
	return a
}

func (s *Stack) Top() int {
	return s.d[s.t]
}
func (s *Stack) IsEmpty() bool {
	return s.t < 0
}

func largestRectangleArea(heights []int) int {
	l := len(heights)
	stack := NewStack()
	mx := 0
	stack.Push(-1)
	for i := 0; i < l; i++ {
		t := stack.Top() // Right Bound
		for stack.Top() != -1 && heights[stack.Top()] > heights[i] {
			h := heights[stack.Pop()]
			l := stack.Top() // Left Bound
			mx = max(mx, h*(t-l))
		}
		stack.Push(i)
	}
	for stack.Top() != -1 {
		h := heights[stack.Pop()]
		t := stack.Top()
		mx = max(mx, h*((l-1)-t))
	}
	return mx
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	// fmt.Println(largestRectangleArea([]int{2, 1}))
	fmt.Println(largestRectangleArea([]int{3, 5, 5, 5, 5, 5}))
	fmt.Println(largestRectangleArea([]int{3, 4, 4, 4, 4}))
	fmt.Println(largestRectangleArea([]int{4, 4, 4, 1}))
	fmt.Println(largestRectangleArea([]int{2, 0, 2}))
	fmt.Println(largestRectangleArea([]int{1, 2, 3, 4, 5}))
	fmt.Println(largestRectangleArea([]int{1}))
	fmt.Println(largestRectangleArea([]int{}))
	fmt.Println(largestRectangleArea([]int{6, 7, 5, 2, 4, 5, 9, 3}))
}
