//https://leetcode.com/problems/implement-queue-using-stacks/

package main

import (
	"errors"
	"fmt"
)

type StackNode struct {
	val  int
	next *StackNode
}

type Stack struct {
	top *StackNode
}

func (s *Stack) Push(item int) {
	s.top = &StackNode{val: item, next: s.top}
}
func (s *Stack) Pop() int {
	if s.top == nil {
		panic(errors.New("stack is empty"))
	}
	v := s.top.val
	s.top = s.top.next
	return v
}

func (s *Stack) Peek() int {
	if s.top != nil {
		return s.top.val
	}
	panic(errors.New("is Empty"))
}
func (s *Stack) isEmpty() bool {
	return s.top == nil
}

type MyQueue struct {
	s1 *Stack
	s2 *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{s1: &Stack{}, s2: &Stack{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

func (this *MyQueue) drain() {
	if this.s2.isEmpty() {
		for !this.s1.isEmpty() {
			this.s2.Push(this.s1.Pop())
		}
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.drain()
	return this.s2.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.drain()
	return this.s2.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.s1.isEmpty() && this.s2.isEmpty()
}

func main() {
	q := Constructor()
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	q.Push(6)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	q.Push(7)
	q.Push(8)
	q.Push(9)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
