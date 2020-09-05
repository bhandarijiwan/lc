// https://leetcode.com/problems/the-skyline-problem/

package main

import (
	"fmt"
	"sort"
)

func Overlap(a, b []int) bool {
	if b[0] > a[1] {
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getSkyline(buildings [][]int) [][]int {
	// height at the start and end of the rectangle
	points := make([][]int, len(buildings)*2)
	for i, j := 0, 0; i < len(buildings); i, j = i+1, j+2 {
		points[j] = []int{buildings[i][0], 0}
		points[j+1] = []int{buildings[i][1], 0}
	}
	for i := 0; i < len(buildings); i++ {
		for j := 0; j < len(points); j++ {
			if points[j][0] >= buildings[i][0] && points[j][0] < buildings[i][1] {
				points[j][1] = max(points[j][1], buildings[i][2])
			}
		}
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	result := make([][]int, len(points))
	j := 0
	prev := []int{0, 0}
	for i := 0; i < len(points); i++ {
		if prev[1] == points[i][1] {
			continue
		}
		prev = points[i]
		result[j] = points[i]
		j++
	}
	return result[:j]
}

func main() {
	input := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	fmt.Println(getSkyline(input))
	input = [][]int{{1, 2, 3}}
	fmt.Println(getSkyline(input))
}
