// https://leetcode.com/problems/generate-parentheses/

package main

import (
	"bytes"
	"fmt"
)

type Stack struct {
	data []rune
	top  int
}

const OPENING = byte('(')
const CLOSING = byte(')')

func (s Stack) isEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

func (s Stack) Peek() rune {
	if s.top >= 0 {
		return s.data[s.top]
	}
	panic("Empty stack")
}
func (s *Stack) Pop() rune {
	if s.top >= 0 {
		a := s.data[s.top]
		s.top = s.top - 1
		return a
	}
	panic("Empty stack!!")
}
func (s *Stack) Push(c rune) {
	s.top = s.top + 1
	if len(s.data) <= s.top {
		newData := make([]rune, len(s.data)*2)
		copy(newData, s.data)
		s.data = newData
	}
	s.data[s.top] = c

}

func NewStack() *Stack {
	return &Stack{data: make([]rune, 2), top: -1}
}

func valid(str string) bool {
	s := NewStack()
	for _, c := range str {
		if c == rune(CLOSING) {
			if s.isEmpty() {
				return false
			}
			s.Pop()
		} else {
			s.Push(c)
		}
	}
	if s.isEmpty() {
		return true
	}
	return false
}

func backTrack(str []byte, f, l int, output *[]string) {
	if f == l {
		s := string(str)
		if valid(s) {
			*output = append(*output, s)
		}
	}
	m := make([]bool, 2)
	for i := f; i < len(str); i++ {
		if m[str[i]-40] {
			continue
		}
		if i != f && str[i] == str[f] {
			continue
		}
		str[i], str[f] = str[f], str[i]
		backTrack(str, f+1, l, output)
		str[i], str[f] = str[f], str[i]
		m[str[i]-40] = true
	}
}

func generateParenthesis(n int) []string {
	str := bytes.Repeat([]byte{OPENING}, n)
	str = append(str, bytes.Repeat([]byte{CLOSING}, n)...)
	output := make([]string, 0)
	backTrack(str, 0, n*2, &output)
	return output
}

func main() {
	s := []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}
	fmt.Println(len(s) == len(generateParenthesis(4)))

}
