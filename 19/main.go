// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	listArr := make([]*ListNode, 0)
	node := head
	for node != nil {
		listArr = append(listArr, node)
		node = node.Next
	}
	l := len(listArr)
	if l-n-1 < 0 {
		if l > 1 {
			return listArr[1]
		}
		return nil
	}
	var dest *ListNode
	if l-n+1 < l {
		dest = listArr[l-n+1]
	}
	listArr[l-n-1].Next = dest
	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	leading, trailing := head, head
	i := 0
	for leading != nil {
		if i > n {
			trailing = trailing.Next
		}
		leading = leading.Next
		i++
	}
	if n >= i {
		if head != nil {
			return head.Next
		}
		return nil
	}
	trailing.Next = trailing.Next.Next
	return head
}

func (l *ListNode) String() string {
	var s string
	node := l
	for node != nil {
		s += "," + strconv.Itoa(node.Val)
		node = node.Next
	}
	return s
}

func main() {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	head := removeNthFromEnd(list, 5)
	fmt.Println(head)

	list = &ListNode{Val: 2, Next: &ListNode{Val: 3}}
	head = removeNthFromEnd(list, 2)
	fmt.Println(head)

	list = &ListNode{Val: 2}
	head = removeNthFromEnd(list, 2)
	fmt.Println(head)
}
