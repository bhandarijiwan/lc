// https://leetcode.com/problems/trapping-rain-water/

package main

import (
	"fmt"
)

type D struct {
	v int
	i int
}

type Stack struct {
	d []D
	t int
}

func NewStack() *Stack {
	return &Stack{d: []D{D{}}, t: -1}
}
func (s *Stack) Push(d D) {
	s.t += 1
	if len(s.d) >= s.t {
		n := make([]D, len(s.d)*2)
		copy(n, s.d)
		s.d = n
	}
	s.d[s.t] = d
}
func (s *Stack) Pop() D {
	if s.t >= 0 {
		d := s.d[s.t]
		l := len(s.d) >> 2
		if s.t < l {
			s.d = s.d[0:l]
		}
		s.t -= 1
		return d
	}
	panic("Stack is empty!!")
}

func (s *Stack) Peek() D {
	return s.d[s.t]
}
func (s *Stack) IsEmpty() bool {
	return s.t < 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int {
	l := len(height)
	if l < 1 {
		return 0
	}
	prev := height[0]
	stack := NewStack()
	amount := 0
	for i := 0; i < l; i++ {
		delta := height[i] - prev
		if delta < 0 {
			stack.Push(D{i: i, v: delta * -1})
		} else if delta > 0 {
			for !stack.IsEmpty() {
				p := stack.Pop()
				amount += min(p.v, delta) * (i - p.i)
				if p.v >= delta {
					if p.v > delta {
						p.v = p.v - delta
						stack.Push(p)
					}
					break
				} else {
					delta -= p.v
				}
			}
		}
		prev = height[i]
	}
	return amount
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{}))
	fmt.Println(trap([]int{1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
