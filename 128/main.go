// https://leetcode.com/problems/longest-consecutive-sequence/
package main

import (
	"fmt"
)

type Node struct {
	rank  int
	count int
	p     *Node
}

func makeSet(val int) *Node {
	node := Node{rank: 0, count: 1}
	node.p = &node
	return &node
}

func Union(x, y *Node) {
	Link(FindSet(x), FindSet(y))
}

func Link(x, y *Node) {
	if x.rank > y.rank {
		y.p = x
		x.count += y.count
	} else {
		x.p = y
		y.count += x.count
		if x.rank == y.rank {
			y.rank = y.rank + 1
		}
	}
}

func FindSet(x *Node) *Node {
	if x != x.p {
		x.p = FindSet(x.p)
	}
	return x.p
}

func longestConsecutive(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	indexMap := make(map[int]int)
	nodes := make([]*Node, len(nums))
	for index, num := range nums {
		if _, ok := indexMap[num]; ok {
			continue
		}
		indexMap[num] = index
		nodes[index] = makeSet(num)
	}
	maxCount := 1
	for i, num := range nums {
		index, _ := indexMap[num]
		if i != index {
			continue
		}
		minusOne, plusOne := num-1, num+1
		plusOneIndex, plusOneExists := indexMap[plusOne]
		if plusOneExists && plusOneIndex > index {
			Union(nodes[index], nodes[plusOneIndex])
		}
		minusOneIndex, minusOneExists := indexMap[minusOne]
		if minusOneExists && minusOneIndex > index {
			Union(nodes[index], nodes[minusOneIndex])
		}
		if nodes[index].p.count > maxCount {
			maxCount = nodes[index].p.count
		}
	}
	return maxCount
}

func main() {
	input := []int{100, 4, 200, 3, 2, 1}
	input = []int{0, 0}
	input = []int{-2, -3, -3, 7, -3, 0, 5, 0, -8, -4, -1, 2}
	fmt.Println(longestConsecutive(input))
}
