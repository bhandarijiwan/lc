// https://leetcode.com/problems/target-sum/
/**
#backtracking
*/

package main

import (
	"fmt"
)

func sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

func helper(nums []int, targetSum int, sum int, i int) int {
	if i >= len(nums) {
		return 0
	}
	count := 0
	for ; i < len(nums); i++ {
		s := (sum - nums[i]) - nums[i]
		if s == targetSum {
			count++
		}
		if s >= targetSum {
			count += helper(nums, targetSum, s, i+1)
		}
	}
	return count
}

func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	count := 0
	if sum == S {
		count++
	}
	return count + helper(nums, S, sum, 0)
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
