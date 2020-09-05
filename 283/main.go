//https://leetcode.com/problems/move-zeroes/
package main

import "fmt"

func moveZeroes(nums []int) {
	i, j, l := 0, 0, len(nums)
	for j < l {
		for j < l && nums[j] != 0 {
			j++
		}
		for i = j; i < l && nums[i] == 0; i++ {
		}
		if i < l {
			a := nums[j]
			nums[j] = nums[i]
			nums[i] = a
		}
		j++
	}
}

func main() {
	input := []int{0,1,0}
	moveZeroes(input)
	fmt.Println(input)
}
