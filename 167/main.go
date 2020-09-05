//https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

package main

import (
	"fmt"
	"sort"
)

func twoSum(numbers []int, target int) []int {
	r := make([]int, 2)
	j := 0
	l := len(numbers)
	i := sort.SearchInts(numbers, target)
	if i >= l {
		i = l - 1
	}
	if i+1 < l {
		i++
	}
	for i > j {
		s := numbers[i] + numbers[j]
		if s == target {
			r[0] = j + 1
			r[1] = i + 1
			break
		} else if s > target {
			i--
		} else {
			j++
		}
	}
	return r
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{-1, 0}, -1))
	fmt.Println(twoSum([]int{-3, 3, 4, 90}, 0))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
}
