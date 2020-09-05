// https://leetcode.com/problems/k-closest-points-to-origin/
// // https://leetcode.com/problems/k-closest-points-to-origin/discuss/220235/Java-Three-solutions-to-this-classical-K-th-problem./270776
package main

import (
	"container/heap"
	"fmt"
)

type E struct {
	d int
	i int
}
type MinHeap struct {
	elem []E
	t    int
	c    int
}

func NewHeap() *MinHeap {
	return &MinHeap{elem: make([]E, 0)}
}

func (h *MinHeap) Swap(i, j int) {
	a := h.elem[i]
	h.elem[i] = h.elem[j]
	h.elem[j] = a
}

func (h *MinHeap) Add(n int, i int) {
	h.elem = append(h.elem, E{n, i})
	h.Swim()

}

func (h *MinHeap) Remove() int {
	l := len(h.elem)
	if l > 0 {
		a := h.elem[0]
		h.Swap(0, l-1)
		h.elem = h.elem[0 : l-1]
		h.Sink()
		return a.i
	}
	panic("queue is empty")
}

func (h *MinHeap) Swim() {
	k := len(h.elem) - 1
	for k > 0 {
		p := k / 2
		if k%2 == 0 {
			p = (k - 1) / 2
		}
		if h.elem[k].d > h.elem[p].d {
			break
		}
		h.Swap(k, p)
		k = p
	}
}
func (h *MinHeap) Sink() {
	l := len(h.elem)
	k := 0
	for k < l {
		left := (2 * k) + 1
		right := (2 * k) + 2
		if right < l {
			v := h.elem[k]
			lv := h.elem[left]
			if h.elem[right].d < lv.d && h.elem[right].d < v.d {
				h.Swap(k, right)
				k = right
			} else if lv.d < v.d {
				h.Swap(k, left)
				k = left
			} else {
				break
			}
		} else if left < l {
			if h.elem[left].d < h.elem[k].d {
				h.Swap(k, left)
				k = left
			} else {
				break
			}
		} else {
			break
		}
	}
}

func distFromOrigin(p1 []int) int {
	return (p1[0] * p1[0]) + (p1[1] * p1[1])
}

func kClosest1(points [][]int, K int) [][]int {
	l := len(points)
	if l <= K {
		return points
	}
	r := make([][]int, K)
	h := NewHeap()
	for i := 0; i < len(points); i++ {
		d := distFromOrigin(points[i])
		h.Add(d, i)
	}
	for i := 0; i < K; i++ {
		idx := h.Remove()
		r[i] = points[idx]
	}
	return r
}

type MinHeapStd []int

func (h MinHeapStd) Less(i, j int) bool {
	return h[j] > h[i]
}
func (h MinHeapStd) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeapStd) Len() int {
	return len(h)
}
func (h *MinHeapStd) Push(item interface{}) {
	x := item.(int)
	*h = append(*h, x)
}
func (h *MinHeapStd) Pop() interface{} {
	old := *h
	l := len(old)
	item := old[l-1]
	*h = old[0 : l-1]
	return item
}

// func kClosest(points [][]int, K int) [][]int {
//
// }

func main() {
	// fmt.Println(kClosest1([][]int{{1, 3}, {-2, 2}}, 1))
	// fmt.Println(kClosest1([][]int{{5, 5}, {5, 7}}, 3))
	i := &MinHeapStd{10, 23, 3, 1, 9}
	heap.Init(i)
	fmt.Println(heap.Pop(i))
	fmt.Println(heap.Pop(i))
	fmt.Println(heap.Pop(i))
	fmt.Println(heap.Pop(i))
}
