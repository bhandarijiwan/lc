// https://leetcode.com/problems/combination-sum/

package main

import (
	"fmt"
)

func fill(a, n int) []int {
	o := make([]int, n)
	for i := 0; i < n; i++ {
		o[i] = a
	}
	return o
}

func combinationSum1(candidates []int, target int) [][]int {
	output := make([][]int, 0)
	for _, candidate := range candidates {
		k := target / candidate
		if target%candidate == 0 {
			output = append(output, fill(candidate, k))
		}

	}
	return output
}

func backTrack(list *[][]int, tempList *[]int, nums []int, target int, start int) {
	if target < 0 {
		return
	}
	if target == 0 {
		a := make([]int, len(*tempList))
		copy(a, *tempList)
		*list = append(*list, a)
		return
	}
	for i := start; i < len(nums); i++ {
		*tempList = append(*tempList, nums[i])
		backTrack(list, tempList, nums, target-nums[i], i)
		t := *tempList
		*tempList = t[0 : len(t)-1]
	}

}

func combinationSum(candidates []int, target int) [][]int {
	output := make([][]int, 0)
	tempList := make([]int, 0)
	backTrack(&output, &tempList, candidates, target, 0)
	return output
}

func main() {
	input := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(input, target))
}
