// https://leetcode.com/problems/count-of-smaller-numbers-after-self/
// #binary indexed tree
package main

import (
	"fmt"
	"sort"
)

type Bt struct {
	d []int
	n int
}

func NewBt(n int) *Bt {
	return &Bt{d: make([]int, n+1), n: n + 1}
}
func (bt *Bt) Add(i, v int) {
	j := i
	for j < bt.n {
		bt.d[j] += v
		j += j & -j
	}
}
func (bt *Bt) Sum(i int) int {
	res := 0
	j := i
	for j > 0 {
		res += bt.d[j]
		j -= j & -j
	}
	return res
}

func countSmaller(nums []int) []int {
	l := len(nums)
	s := make([]int, l)
	copy(s, nums)
	sort.Ints(s)
	r := make(map[int]int)
	for i, i2 := range s {
		r[i2] = i + 1
	}
	res := make([]int, l)
	bt := NewBt(l)
	for i := l - 1; i >= 0; i-- {
		res[i] = bt.Sum(r[nums[i]]-1)
		bt.Add(r[nums[i]], 1)
	}
	return res
}

func main() {
	input := []int{5, 2, 6, 1}
	fmt.Println(countSmaller(input))
}
