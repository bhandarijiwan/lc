// https://leetcode.com/problems/meeting-scheduler/

package main

import (
	"container/heap"
	"fmt"
)

type PQ [][]int

func (a PQ) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}
func (a PQ) Len() int {
	return len(a)
}
func (a PQ) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a *PQ) Push(item interface{}) {
	*a = append(*a, item.([]int))
}
func (a *PQ) Pop() interface{} {
	old := *a
	item := old[len(old)-1]
	*a = old[0 : len(old)-1]
	return item
}

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	pq := make(PQ, 0)
	for _, s := range slots1 {
		pq.Push(s)
	}
	for _, s := range slots2 {
		pq.Push(s)
	}
	heap.Init(&pq)
	for {
		item := heap.Pop(&pq).([]int)
		if len(pq) < 1 {
			break
		}
		top := pq[0]
		if item[1] >= top[0]+duration && top[1]-top[0] >= duration {
			return []int{top[0], top[0] + duration}
		}
	}
	return []int{}
}

func main() {
	s1 := [][]int{{10, 50}, {60, 120}, {140, 210}}
	s2 := [][]int{{0, 15}, {60, 70}}
	s1 = [][]int{{216397070, 363167701}, {98730764, 158208909}, {441003187, 466254040}, {558239978, 678368334}, {683942980, 717766451}}
	s2 = [][]int{{50490609, 222653186}, {512711631, 670791418}, {730229023, 802410205}, {812553104, 891266775}, {230032010, 399152578}}
	fmt.Println(minAvailableDuration(s1, s2, 1))
}
