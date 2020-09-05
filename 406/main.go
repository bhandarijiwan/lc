// https://leetcode.com/problems/queue-reconstruction-by-height/

package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	output := make([][]int, len(people))
	l := len(output)
	for _, p := range people {
		prev := output[p[1]]
		for i := p[1]; i < l; i++ {
			output[i], prev = prev, output[i]
		}
		output[p[1]] = p
	}
	return output
}

func main() {
	input := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(input))
}
