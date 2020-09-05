// https://leetcode.com/problems/counting-bits/

package main

import (
	"fmt"
	"math/bits"
)

func countBits(num int) []int {
	s := make([]int, num+1)
	for i := uint(0); i <= uint(num); i++ {
		s[i] = bits.OnesCount(i)
	}
	return s
}

func main() {
	fmt.Println(countBits(5))
}
