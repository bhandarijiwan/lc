// https://leetcode.com/problems/partition-equal-subset-sum/
// #0/1 Knapsack

package main

import (
	"fmt"
)

func KnapSack(v []int, w []int, W int) {
	V := make([][]int, len(v))
	for i := 0; i < len(V); i++ {
		V[i] = make([]int, W+1)
	}
	for j := 1; j <= W; j++ {
		if j >= w[0] {
			V[0][j] = v[0]
		}
	}
	for i := 1; i < len(v); i++ {
		for j := 1; j <= W; j++ {
			if j >= w[i] {
				V[i][j] = V[i-1][j]
				// take current
				s := v[i] + V[i-1][j-w[i]]
				if s > V[i][j] {
					V[i][j] = s
				}
			} else {
				V[i][j] = V[i-1][j]
			}
		}
	}
	fmt.Println(V)
}

func canPartition1(nums []int) bool {
	l := len(nums)
	sum := 0
	for i := 0; i < l; i++ {
		sum += nums[i]
	}
	if sum%2 > 0 {
		return false
	}
	target := sum / 2
	r := []int{0}
	for i := 0; i < len(nums); i++ {
		for _, i2 := range r {
			r = append(r, i2+nums[i])
		}
	}
	for _, i2 := range r {
		if i2 == target {
			return true
		}
	}
	return false
}

func canPartition2(nums []int) bool {
	l := len(nums)
	sum := 0
	for i := 0; i < l; i++ {
		sum += nums[i]
	}
	if sum%2 > 0 {
		return false
	}
	target := sum / 2
	V := make([][]bool, len(nums))
	for i := 0; i < len(V); i++ {
		V[i] = make([]bool, target+1)
	}
	for j := 1; j <= target; j++ {
		if nums[0] == j {
			V[0][j] = true
		}
	}
	for i := 1; i < l; i++ {
		for j := 1; j <= target; j++ {
			if j >= nums[i] {
				V[i][j] = V[i-1][j] || V[i-1][j-nums[i]]
			} else {
				V[i][j] = V[i-1][j]
			}
		}
	}
	return V[l-1][target]
}
func canPartition(nums []int) bool {
	l := len(nums)
	sum := 0
	for i := 0; i < l; i++ {
		sum += nums[i]
	}
	if sum%2 > 0 {
		return false
	}
	target := sum / 2
	V := make([]bool, target+1)
	for j := 1; j <= target; j++ {
		if nums[0] == j {
			V[j] = true
		}
	}
	// from left to right means dp[i][j] = dp[i][j-nums[i-1]])
	// form right to left means dp[i][j] = dp[i-1][j-nums[i-1]]
	for i := 1; i < l; i++ {
		for j := target; j >= 1; j-- {
			if j >= nums[i] {
				V[j] = V[j] || V[j-nums[i]]
			}
		}
	}
	return V[target]
}

func main() {
	fmt.Println(canPartition([]int{100}))
	fmt.Println(canPartition([]int{2, 2, 3, 5}))
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
	fmt.Println(canPartition([]int{1, 1}))
	fmt.Println(canPartition([]int{1, 1, 2, 2}))
	fmt.Println(canPartition([]int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}))

	// w := []int{5, 4, 6, 3}
	// v := []int{10, 40, 30, 50}
	// KnapSack(v, w, 10)
}
