// https://leetcode.com/problems/find-the-duplicate-number/

package main

import (
	"fmt"
)

func findDuplicate(nums []int) int {
	// treat nums as a linked list
	// First find the intersection points of slow and fast pointer
	slowPtr, fastPtr := 0, 0
	slow := nums[slowPtr]
	fast := -slow
	for slow != fast {
		slowPtr, fastPtr = nums[slowPtr], nums[nums[fastPtr]]
		slow, fast = nums[slowPtr], nums[fastPtr]
	}

	// Find the start of loop
	slowPtr = 0
	slow = nums[slowPtr]
	for slow != fast {
		slowPtr, fastPtr = nums[slowPtr], nums[fastPtr]
		slow, fast = nums[slowPtr], nums[fastPtr]
	}
	return slow
}

func main() {
	fmt.Println(findDuplicate([]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}))
}
