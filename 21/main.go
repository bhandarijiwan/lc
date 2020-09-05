//https://leetcode.com/problems/merge-two-sorted-lists/

package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	n := l
	s := ""
	for n != nil {
		s = s + strconv.Itoa(n.Val) + ","
		n = n.Next
	}
	return s
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l := &ListNode{}
	n1 := l1
	n2 := l2
	n := l
	var prev *ListNode
	for n1 != nil && n2 != nil {
		if n == nil {
			n = &ListNode{}
		}
		if prev != nil {
			prev.Next = n
		}
		if n1.Val == n2.Val {
			n.Val = n1.Val
			n1 = n1.Next
		} else if n1.Val < n2.Val {
			n.Val = n1.Val
			n1 = n1.Next
		} else {
			n.Val = n2.Val
			n2 = n2.Next
		}
		prev = n
		n = n.Next
	}
	for n1 != nil {
		p := &ListNode{Val: n1.Val}
		prev.Next = p
		prev = p
		n1 = n1.Next
	}
	for n2 != nil {
		p := &ListNode{Val: n2.Val}
		prev.Next = p
		prev = p
		n2 = n2.Next
	}
	return l
}

func main() {
	//l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	//l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	fmt.Println(mergeTwoLists(&ListNode{Val: 1}, nil))
}
