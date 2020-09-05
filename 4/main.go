// https://leetcode.com/problems/median-of-two-sorted-arrays/

package main

import (
	"fmt"
)

func checkOverlap(nums1 []int, nums2 []int) bool {
	l1, l2 := len(nums1), len(nums2)
	if l1 < 1 || l2 < 1 {
		return false
	}
	if nums1[0] >= nums2[l2-1] {
		return false
	}
	if nums2[0] >= nums1[l1-1] {
		return false
	}
	return true
}

func merge(nums1 []int, nums2 []int) []int {
	l1, l2 := len(nums1), len(nums2)
	m := make([]int, l1+l2)
	i, j, k := 0, 0, 0
	for i < l1 && j < l2 {
		if nums1[i] <= nums2[j] {
			m[k] = nums1[i]
			i++
		} else {
			m[k] = nums2[j]
			j++
		}
		k++
	}
	for i < l1 {
		m[k] = nums1[i]
		i++
		k++
	}
	for j < l2 {
		m[k] = nums2[j]
		j++
		k++
	}
	return m
}

func median(m []int) float64 {
	l := len(m)
	if l%2 == 0 {
		i := l >> 1
		return (float64(m[i-1]) + float64(m[i])) / 2
	}
	return float64(m[l>>1])
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := merge(nums1, nums2)
	return median(m)
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	m := findMedianSortedArrays(nums1, nums2)
	fmt.Println(m)
	fmt.Println(findMedianSortedArrays([]int{}, []int{1, 3, 4}))
}
