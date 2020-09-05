//https://leetcode.com/problems/kth-largest-element-in-a-stream/

package main

import (
	"fmt"
	"sort"
)

type MinHeap struct {
	s []int
}

func (h *MinHeap) Add(n int) {
	l := len(h.s)
	h.s[l-1] = n
	h.heapify()
}

func (h *MinHeap) Pop() int {
	l := len(h.s)
	t := h.s[l-1]
	h.s[l-1] = h.s[0]
	h.s[0] = t
	a := h.s[l-1]
	h.heapify()
	return a
}

// Restore the heap property
func (h *MinHeap) heapify() {
	l := len(h.s) - 1
	for l >= 0 {
		p := (l + 1) / 2
		if h.s[l] > h.s[p] {
			break
		}
		a := h.s[l]
		h.s[l] = h.s[p]
		h.s[p] = a
		l = p
	}
}

func initHeap(k int, nums []int) *MinHeap {
	sort.Ints(nums)
	return &MinHeap{s: nums[0:k]}
}

type KthLargest struct {
	h *MinHeap
	k int
}

func (kt *KthLargest) Add(val int) int {
	k := kt.h.s[0]
	if val <= k {
		return k
	}
	kt.h.Pop()
	kt.h.Add(val)
	return kt.h.s[0]
}

func Constructor(k int, nums []int) KthLargest {
	v := KthLargest{h: initHeap(k, nums), k: k}
	return v
}

func main() {
	v := Constructor(3, []int{4, 5, 8, 2})
	n := []int{3, 5, 10, 9, 4}
	for _, c := range n {
		fmt.Println(v.Add(c))
	}
}
