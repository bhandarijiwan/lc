// https://leetcode.com/problems/rectangle-overlap/
package main

import (
	"fmt"
)

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	left := rec2[2] <= rec1[0]
	right := rec2[0] >= rec1[2]
	top := rec2[1] >= rec1[3]
	bottom := rec2[3] <= rec1[1]
	return !(left || right || top || bottom)
}

func main() {
	fmt.Println(isRectangleOverlap([]int{-6, -10, 9, 2}, []int{0, 5, 4, 8}))
}
