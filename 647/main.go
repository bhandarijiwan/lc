// https://leetcode.com/problems/palindromic-substrings/

package main

import (
	"fmt"
)

func countSubstrings(s string) int {
	l := len(s)
	count := 0
	for i := 0; i < 2*l-1; i++ {
		left := i / 2
		right := left + i%2
		for left >= 0 && right < l {
			if s[left] != s[right] {
				break
			}
			count = count + 1
			left = left - 1
			right = right + 1
		}
	}
	return count
}

func main() {
	fmt.Println(countSubstrings("dnncbwoneinoplypwgbwktmvkoimcooyiwirgbxlcttgteqthcvyoueyftiwgwwxvxvg"))
	// fmt.Println(countSubstrings("length"))

}
