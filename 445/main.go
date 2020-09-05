// https://leetcode.com/problems/add-two-numbers-ii/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// go:inline
func reverse(l *ListNode) *ListNode {
	var prev *ListNode
	for l != nil {
		n := l.Next
		l.Next = prev
		prev = l
		l = n
	}
	return prev
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	rl1, rl2 := l1, l2
	rl1, rl2 = reverse(rl1), reverse(rl2)
	var prev *ListNode
	h1, h2 := rl1, rl2
	c := 0
	for h1 != nil && h2 != nil {
		s := h1.Val + h2.Val + c
		c = 0
		if s >= 10 {
			c = 1
			s = s - 10
		}
		n := &ListNode{Val: s}
		if prev != nil {
			n.Next = prev
		} else {
			prev = n
		}
		prev = n
		h1 = h1.Next
		h2 = h2.Next
	}
	h3 := h1
	if h2 != nil {
		h3 = h2
	}
	if prev != nil {
		for h3 != nil {
			h3.Val = h3.Val + c
			c = 0
			if h3.Val >= 10 {
				h3.Val = h3.Val - 10
				c = 1
			}
			a := h3.Next
			h3.Next = prev
			prev = h3
			h3 = a
		}
		if c > 0 {
			return &ListNode{Val: c, Next: prev}
		}
		return prev
	}
	return reverse(h3)
}

func main() {
	l1 := &ListNode{Val: 7, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}
	l2 := &ListNode{Val: 7}
	rl1 := addTwoNumbers(l1, l2)
	for rl1 != nil {
		fmt.Println(rl1.Val)
		rl1 = rl1.Next
	}
}
