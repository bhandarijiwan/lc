// https://leetcode.com/problems/linked-list-cycle/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle1(head *ListNode) bool {
	n := head
	s := &ListNode{Val: 0}
	var p *ListNode
	for n != nil {
		if n == s {
			return true
		}
		if p != nil {
			p.Next = s
		}
		p = n
		n = n.Next
	}
	return false
}
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func main() {
	l1 := &ListNode{Val: 3}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 0, Next: &ListNode{Val: 4, Next: nil}}
	l1.Next = l2
	l2.Next = l3
	fmt.Println(hasCycle(l1))
}
