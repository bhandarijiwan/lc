// https://leetcode.com/problems/merge-intervals/

package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func overlaps(a, b []int) bool {
	if b[0] > a[1] || a[0] > b[1] {
		return false
	}
	return true
}

func merge(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] <= intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	j := 0
	for i := 0; i < len(intervals); i++ {
		if overlaps(intervals[j], intervals[i]) {
			intervals[j][1] = max(intervals[j][1], intervals[i][1])
		} else {
			j = j + 1
			intervals[j] = intervals[i]
		}
	}
	return intervals[:j+1]
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {16, 20}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 2}, {3, 5}}))
}
