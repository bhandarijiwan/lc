// https://leetcode.com/problems/sliding-window-maximum/
// Deque

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	v    int
	prev *Node
	next *Node
}
type Deque struct {
	head *Node
	tail *Node
	nums []int
}

func NewDeque() *Deque {
	sentinel := &Node{}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return &Deque{head: sentinel, tail: sentinel}
}

func (d *Deque) String() string {
	n := d.head.next
	sb := strings.Builder{}
	for n != d.head {
		sb.WriteString(strconv.Itoa(d.nums[n.v]) + ",")
		n = n.next
	}
	return sb.String()
}

func (d *Deque) PushHead(item int) {
	node := &Node{v: item, prev: d.head, next: d.head.next}
	d.head.next.prev = node
	d.head.next = node
}

func (d *Deque) IsEmpty() bool {
	return d.head == d.head.next
}
func (d *Deque) GetHead() int {
	return d.head.next.v
}
func (d *Deque) GetTail() int {
	return d.tail.prev.v
}
func (d *Deque) PushTail(item int) {
	node := &Node{v: item, prev: d.tail.prev, next: d.tail}
	d.tail.prev.next = node
	d.tail.prev = node
}
func (d *Deque) RemoveTail() int {
	node := d.tail.prev
	node.prev.next = node.next
	node.next.prev = node.prev
	return node.v
}

func (d *Deque) RemoveHead() int {
	node := d.head.next
	d.head.next = node.next
	node.next.prev = node.prev
	return node.v
}

func maxSlidingWindow(nums []int, k int) []int {
	l := len(nums)
	if l <= 1 || k <= 1 {
		return nums
	}

	d := NewDeque()
	d.nums = nums
	maxIndex := 0
	output := make([]int, 0)
	cleanDeque := func(i, k int) {
		if !d.IsEmpty() && d.GetHead() == i-k {
			d.RemoveHead()
		}
		for !d.IsEmpty() && nums[i] > nums[d.GetTail()] {
			d.RemoveTail()
		}
	}
	for i := 0; i < k; i++ {
		cleanDeque(i, k)
		d.PushTail(i)
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
		fmt.Println(d)
	}
	output = append(output, nums[maxIndex])
	for i := k; i < l; i++ {
		cleanDeque(i, k)
		d.PushTail(i)
		output = append(output, nums[d.GetHead()])
		// fmt.Println(d)
	}
	return output
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	// fmt.Println(maxSlidingWindow([]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4))
	// fmt.Println(maxSlidingWindow([]int{-1, 1}, 1))

}
