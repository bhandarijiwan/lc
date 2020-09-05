// https://leetcode.com/problems/merge-sorted-array/

package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, len(nums1)-1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	if i >= 0 {
		copy(nums1[:k+1], nums1[:i+1])
	}
	if j >= 0 {
		copy(nums1[:k+1], nums2[:j+1])
	}
}

func main() {
	nums1 := []int{0,0,3,0,0,0,0,0,0}
	nums2 := []int{-1,1,1,1,2,3}
	merge(nums1, 3, nums2, 6)
	fmt.Println(nums1)
}
