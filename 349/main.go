// https://leetcode.com/problems/intersection-of-two-arrays/

package main

import (
	"fmt"
	"sort"
)

func intersection1(nums1 []int, nums2 []int) []int {
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	m := make(map[int]bool)
	output := make([]int, 0)
	for _, num := range nums2 {
		if m[num] {
			continue
		}
		r := sort.SearchInts(nums1, num)
		if r < len(nums1) && nums1[r] == num {
			output = append(output, num)
		}
		m[num] = true
	}
	return output
}

func intersection2(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool)
	for _, num := range nums1 {
		m1[num] = true
	}
	output := make([]int, 0)
	m2 := make(map[int]bool)
	for _, num := range nums2 {
		if m2[num] {
			continue
		}
		m2[num] = true
		if m1[num] {
			output = append(output, num)
		}
	}
	return output
}

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	long, short := nums1, nums2
	if len(long) < len(short) {
		long, short = nums2, nums1
	}
	// align shorter with the longer
	indexLong := 0
	for ; indexLong < len(long); indexLong++ {
		if short[0] == long[indexLong] {
			break
		}
	}
	indexShort := 0
	output := make([]int, 0)
	for indexLong < len(long) && indexShort < len(short) {
		if short[indexShort] == long[indexLong] {
			output = append(output, short[indexShort])
		}
		prevIndexLong := indexLong
		for indexLong < len(long) && long[prevIndexLong] == long[indexLong] {
			indexLong++
		}
		prevIndexShort := indexShort
		for indexShort < len(short) && short[prevIndexShort] == short[indexShort] {
			indexShort++
		}
	}
	return output
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
