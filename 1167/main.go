// https://leetcode.com/problems/minimum-cost-to-connect-sticks/

package main

import (
	"container/heap"
	"fmt"
)

type H []int

func (h H) Len() int           { return len(h) }
func (h H) Less(i, j int) bool { return h[i] < h[j] }
func (h H) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *H) Push(e interface{}) {
	*h = append(*h, e.(int))
}
func (h *H) Pop() interface{} {
	old := *h
	item := old[len(old)-1]
	*h = old[:len(old)-1]
	return item
}

func connectSticks(sticks []int) int {
	s := H(sticks)
	heap.Init(&s)
	cost := 0
	for i := 0; i < len(sticks)-1; i++ {
		x, y := heap.Pop(&s).(int), heap.Pop(&s).(int)
		cost += x + y
		heap.Push(&s, x+y)
	}
	return cost
}

func main() {
	fmt.Println(connectSticks([]int{2, 4, 3}))
	fmt.Println(connectSticks([]int{2, 1, 1}))
	fmt.Println(connectSticks([]int{2, 5, 4}))
	fmt.Println(connectSticks([]int{5, 1}))
	fmt.Println(connectSticks([]int{3, 4, 5, 8, 9}))
	fmt.Println(connectSticks([]int{3354, 4316, 3259, 4904, 4598, 474, 3166, 6322, 8080, 9009}))

}
