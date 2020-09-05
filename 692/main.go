// https://leetcode.com/problems/top-k-frequent-words/

package main

import (
	"container/heap"
	"fmt"
)

type Element struct {
	count int
	val   string
}
type PQ []Element

func (p PQ) Less(i, j int) bool {
	if p[i].count == p[j].count {
		return p[i].val < p[j].val
	}
	return p[i].count > p[j].count
}
func (p PQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PQ) Len() int {
	return len(p)
}
func (p *PQ) Push(item interface{}) {
	elemItem := item.(Element)
	*p = append(*p, elemItem)
}
func (p *PQ) Pop() interface{} {
	a := *p
	l := len(a)
	item := a[l-1]
	*p = a[0 : l-1]
	return item.val
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, s := range words {
		m[s]++
	}
	h := make(PQ, 0)
	for str, count := range m {
		h = append(h, Element{count: count, val: str})
	}
	heap.Init(&h)
	output := make([]string, k)
	for i := 0; i < k; i++ {
		output[i] = heap.Pop(&h).(string)
	}
	return output
}

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
}
