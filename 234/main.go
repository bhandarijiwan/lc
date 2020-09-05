// https://leetcode.com/problems/palindrome-linked-list/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	c := 0
	n := head
	for n != nil {
		c = c + 1
		n = n.Next
	}
	i, m := 0, c/2
	n = head
	var p *ListNode
	for i < m {
		a := n.Next
		n.Next = p
		p = n
		n = a
		i = i + 1
	}
	if c%2 > 0 {
		n = n.Next
	}
	for p != nil && n != nil {
		if p.Val != n.Val {
			return false
		}
		p = p.Next
		n = n.Next
	}
	return true
}

func main() {

	fmt.Print(isPalindrome(&ListNode{Val: 1}))
}
