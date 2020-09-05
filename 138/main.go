// https://leetcode.com/problems/copy-list-with-random-pointer/

package main

import (
	"fmt"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList1(head *Node) *Node {
	var prev *Node
	m := make(map[*Node]int)
	arr := make([]*Node, 0)
	n := head
	i := 0
	for n != nil {
		newNode := &Node{Val: n.Val}
		if prev != nil {
			prev.Next = newNode
		} else {
			prev = newNode
		}
		m[n] = i
		arr = append(arr, newNode)
		prev = newNode
		n = n.Next
		i += 1
	}
	n = head
	i = 0
	for n != nil {
		if n.Random != nil {
			index, _ := m[n.Random]
			arr[i].Random = arr[index]
		}
		n = n.Next
		i += 1
	}
	if len(arr) > 0 {
		return arr[0]
	}
	return nil
}

func copyRandomList(head *Node) *Node {
	// interleaved list
	node := head
	for node != nil {
		newNode := &Node{Val: node.Val}
		newNode.Next = node.Next
		node.Next = newNode
		node = newNode.Next
	}
	node = head
	// var newHead *Node
	// var prev *Node
	for node != nil {
		oldRandom := node.Random
		if oldRandom != nil {
			node.Next.Random = oldRandom.Next
		}
		// if prev != nil {
		// 	prev.Next = node.Next
		// } else {
		// 	newHead = node.Next
		// }
		// prev = node.Next
		// node.Next = node.Next.Next
		// node = node.Next
		node = node.Next.Next

	}
	node = head.Next
	printList(head.Next, head.Next)
	return head
}

func printList(n *Node, old *Node) {
	head := n
	oldHead := old
	for head != nil {
		fmt.Print(head.Val, "(")
		if head.Random != nil {
			fmt.Print(head.Random.Val)
		} else {
			fmt.Print(head.Random)
		}
		fmt.Print(" ", head.Random == oldHead.Random, "), ")
		head = head.Next
		oldHead = oldHead.Next
	}
	fmt.Println()
}
func main() {
	input := []*Node{&Node{Val: 7}, &Node{Val: 13}, &Node{Val: 11}, &Node{Val: 10}, &Node{Val: 1}}
	input[0].Next = input[1]
	input[1].Next = input[2]
	input[2].Next = input[3]
	input[3].Next = input[4]
	input[1].Random = input[0]
	input[2].Random = input[4]
	input[3].Random = input[2]
	input[4].Random = input[0]

	output := copyRandomList(input[0])
	printList(output, input[0])

	// output = copyRandomList(nil)
	// printList(output, nil)

	node := &Node{Val: 90}
	node.Random = node

	output = copyRandomList(node)
	printList(output, node)
}
