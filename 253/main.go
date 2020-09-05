// https://leetcode.com/problems/meeting-rooms-ii/

package main

import (
	"fmt"
	"sort"
)

func doesOverlap(a []int, b []int) bool {
	if a[1] <= b[0] || b[1] <= a[0] {
		return false
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// r can be a min-heap on end-time
	r := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		var j int
		for j = 0; j < len(r); j++ {
			if !doesOverlap(intervals[i], r[j]) {
				break
			}
		}
		if j == len(r) {
			r = append(r, intervals[i])
		} else {
			r[j][0] = min(r[j][0], intervals[i][0])
			r[j][1] = max(r[j][1], intervals[i][1])
		}
	}
	return len(r)
}

func main() {
	input := [][]int{{0, 30}, {1, 4}, {5, 10}, {15, 20}}
	fmt.Println(minMeetingRooms(input))
	input = [][]int{{7, 10}, {2, 4}}
	fmt.Println(minMeetingRooms(input))
	input = [][]int{{16, 21}, {25, 30}, {17, 19}, {19, 20}, {26, 26}, {8, 9}, {15, 23}, {5, 8}, {6, 10}, {0, 3}}
	fmt.Println(minMeetingRooms(input))
	input = [][]int{{13, 15}, {1, 13}}
	fmt.Println(minMeetingRooms(input))
}
