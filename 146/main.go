// https://leetcode.com/problems/lru-cache/
package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	prev *ListNode
	next *ListNode
	key  int
	val  int
}

type DoublyLL struct {
	head     *ListNode
	tail     *ListNode
	capacity int
}

func (ll *DoublyLL) String() string {
	var s string
	n := ll.head.next
	for n != ll.head {
		s += "key=" + strconv.Itoa(n.key) + " val=" + strconv.Itoa(n.val)
		n = n.next
	}
	return s
}

func (ll *DoublyLL) Add(node *ListNode) {
	next := ll.head.next
	node.next = next
	node.prev = ll.head
	next.prev = node
	ll.head.next = node
}
func (ll *DoublyLL) Remove(node *ListNode) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev

}
func (ll *DoublyLL) RemoveTail() *ListNode {
	tail := ll.tail.prev
	tail.prev.next = ll.tail
	ll.tail.prev = tail.prev
	return tail
}
func (ll *DoublyLL) MoveToHead(node *ListNode) {
	ll.Remove(node)
	ll.Add(node)
}

func NewDoublyLL() *DoublyLL {
	// sentinelTail := ListNode{}
	// sentinelHead := ListNode{}
	// sentinelHead.next = &sentinelTail
	// sentinelTail.prev = &sentinelHead
	// return &DoublyLL{head: &sentinelHead, tail: &sentinelTail}
	sentinel := ListNode{}
	ll := &DoublyLL{head: &sentinel, tail: &sentinel}
	ll.head.next = ll.tail
	ll.tail.prev = ll.head
	return ll
}

type LRUCache struct {
	ll   *DoublyLL
	lt   map[int]*ListNode
	cap  int
	size int
}

func Constructor(capacity int) LRUCache {
	lt := make(map[int]*ListNode)
	ll := NewDoublyLL()
	return LRUCache{ll: ll, lt: lt, size: 0, cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.lt[key]
	if !ok {
		return -1
	}
	this.ll.MoveToHead(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.lt[key]
	if !ok {
		newNode := &ListNode{val: value, key: key}
		this.ll.Add(newNode)
		this.lt[key] = newNode
		this.size += 1
		if this.size > this.cap {
			tail := this.ll.RemoveTail()
			delete(this.lt, tail.key)
			this.size -= 1
		}
	} else {
		node.val = value
		this.ll.MoveToHead(node)
	}
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(4))
	fmt.Println(cache.Get(3))
	cache.Put(1, 1)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(1, 1)
	fmt.Println(cache.ll)
}
