// https://leetcode.com/problems/subarray-sum-equals-k/
// #continuous-subarray, #hashtable

package main

import (
	"fmt"
)

func subarraySum1(nums []int, k int) int {
	count := 0
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	p, l := 0, len(nums)
	for i := 0; i < l; i++ {
		s = s - p
		sum, prev := s, 0
		for j := l - 1; j >= i; j-- {
			sum = sum - prev
			if sum == k {
				count++
			}
			prev = nums[j]
		}
		p = nums[i]
	}
	return count
}

func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	count := 0
	sum := 0
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		count += m[sum-k]
		m[sum]++
	}
	return count
}

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 1, 1, 3, -1}, 2))
	fmt.Println(subarraySum([]int{2, 3, 1, -2}, 2))
}
