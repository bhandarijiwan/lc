//https://leetcode.com/problems/reverse-linked-list/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	s := ""
	h := l
	i := 0
	for h != nil {
		t := strings.Repeat("__", i)
		s = s + "\n" + t + "{val: " + strconv.Itoa(h.Val) + ", next: "
		i++
		h = h.Next
		if h == nil {
			s = s + strings.Repeat("}", i)
		}
	}
	return s
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	next := head
	for next != nil {
		cur := next
		next = cur.Next
		cur.Next = prev
		prev = cur
	}
	return prev
}

func main() {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	fmt.Println(reverseList(l))
}
