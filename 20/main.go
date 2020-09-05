// https://leetcode.com/problems/valid-parentheses/

package main

import (
	"fmt"
)

type Stack struct {
	s []byte
	t int
}

func (s *Stack) Push(b byte) {
	s.t++
	s.s[s.t] = b
}

func (s *Stack) Pop() byte {
	if s.t >= 0 {
		b := s.s[s.t]
		s.t--
		return b
	}
	return '-'
}

func NewStack(l int) *Stack {
	return &Stack{s: make([]byte, l), t: -1}
}

func isValid(s string) bool {
	l := len(s)
	stack := NewStack(l)
	for i := 0; i < l; i++ {
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			stack.Push(s[i])
		} else {
			b := stack.Pop()
			if s[i]-b > 2 {
				return false
			}
		}
	}
	if stack.Pop() == '-' {
		return true
	}
	return false
}

func main() {
	fmt.Println(isValid("[(([))]]"))
}
