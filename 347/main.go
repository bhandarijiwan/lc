// https://leetcode.com/problems/top-k-frequent-elements/

package main

import (
	"container/heap"
	"fmt"
)

type E struct {
	v int
	c int
}

type PriorityQueue []E

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].c > pq[j].c
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]

}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(E)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, n := range nums {
		countMap[n] = countMap[n] + 1
	}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for k, v := range countMap {
		heap.Push(&pq, E{c: v, v: k})
	}
	r := make([]int, k)
	for i := 0; i < k; i++ {
		r[i] = heap.Pop(&pq).(E).v
	}
	return r
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 3))
}
