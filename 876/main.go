//https://leetcode.com/problems/middle-of-the-linked-list/
package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	c := 0
	n := head
	for n != nil {
		c = c + 1
		n = n.Next
	}
	m := int(math.Floor(float64(c / 2.0)))
	n = head
	c = 0
	for c < m && n != nil {
		n = n.Next
		c = c + 1
	}
	return n
}

func main() {
	//ll := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 5, Next: nil}}}}}}
	ll1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	fmt.Println(middleNode(ll1))
}
