// https://leetcode.com/problems/decode-string/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	d []string
	t int
}

func NewStack() *Stack {
	return &Stack{d: make([]string, 1), t: -1}
}

func (s *Stack) Push(str string) {
	s.t = s.t + 1
	if s.t >= len(s.d) {
		nd := make([]string, len(s.d)*2)
		copy(nd, s.d)
		s.d = nd
	}
	s.d[s.t] = str
}
func (s *Stack) Pop() string {
	if s.t >= 0 {
		str := s.d[s.t]
		s.t = s.t - 1
		return str
	}
	panic("Empty!!")
}
func (s *Stack) Peek() string {
	return s.d[s.t]
}
func (s *Stack) IsEmpty() bool {
	return s.t < 0
}

func decodeString(s string) string {
	stack := NewStack()
	numStr := make([]rune, 0)
	for _, c := range s {
		if c >= '0' && c <= '9' {
			numStr = append(numStr, c)
			continue
		} else {
			if len(numStr) > 0 {
				stack.Push(string(numStr))
			}
			numStr = numStr[0:0]
		}
		str := string(c)
		if c == ']' {
			str = ""
			for ic := ""; ic != "["; ic = stack.Pop() {
				str = ic + str
			}
			r, _ := strconv.Atoi(stack.Pop())
			str = strings.Repeat(str, r)
		}
		stack.Push(str)
	}
	str := ""
	for !stack.IsEmpty() {
		str = stack.Pop() + str
	}
	return str
}

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef") == "abcabccdcdcdef")
	fmt.Println(decodeString("01[leetcode]12[rocks!]"))
}
