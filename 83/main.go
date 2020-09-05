//https://leetcode.com/problems/remove-duplicates-from-sorted-list/

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
	s := ""
	n := l
	for n != nil {
		s = s + strconv.Itoa(n.Val) + ","
		n = n.Next
	}
	return s
}

func deleteDuplicates(head *ListNode) *ListNode {
	n := head
	for n != nil {
		if n.Next != nil && n.Next.Val == n.Val {
			n.Next = n.Next.Next
		} else {
			n = n.Next
		}
	}
	return head
}

func main() {
	//ll := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{2, &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}
	fmt.Println(deleteDuplicates(&ListNode{Val: 0}))
}
