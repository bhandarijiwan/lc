//https://leetcode.com/problems/split-a-string-in-balanced-strings/

package main

import "fmt"

func balancedStringSplit(s string) int {
	n := len(s)
	count := 0
	rCount := 0
	lCount := 0
	for i := 0; i < n; i = i + 2 {
		c1 := s[i]
		c2 := s[i+1]
		if c1 == 'R' {
			rCount = rCount + 1
		} else {
			lCount = lCount + 1
		}
		if c2 == 'R' {
			rCount = rCount + 1
		} else {
			lCount = lCount + 1
		}
		if rCount == lCount {
			count = count + 1
			rCount = 0
			lCount = 0
		}
	}
	return count
}

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
	fmt.Println(balancedStringSplit("RLLLLRRRLR"))
	fmt.Println(balancedStringSplit("LLLLRRRR"))
	fmt.Println(balancedStringSplit("RLRRRLLRLL"))
}
