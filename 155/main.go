//https://leetcode.com/problems/min-stack/

package main

import "fmt"

type StackNode struct {
	elem int
	min  int
	Next *StackNode
}

type MinStack struct {
	top *StackNode
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{top: nil}
}

func (this *MinStack) Push(x int) {
	n := &StackNode{elem: x, min: x}
	if this.top != nil {
		if this.top.min < x {
			n.min = this.top.min
		}
	}
	n.Next = this.top
	this.top = n
}

func (this *MinStack) Pop() {
	if this.top != nil {
		this.top = this.top.Next
	}
}

func (this *MinStack) Top() int {
	if this.top != nil {
		return this.top.elem
	}
	panic("empty stack")
}

func (this *MinStack) GetMin() int {
	return this.top.min
}

func main() {
	m := Constructor()
	m.Push(-2)
	m.Push(0)
	m.Push(-3)
	fmt.Println(m.Top())
	fmt.Println(m.GetMin())
}
