// https://leetcode.com/problems/valid-palindrome/

package main

import (
	"fmt"
)

const d = 'a' - 'A'

func isLetter(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	if b >= 'a' && b <= 'z' {
		return true
	}
	if b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}

func lower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + d
	}
	return b
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for j > i {
		for i < j && !isLetter(s[i]) {
			i++
		}
		for j > i && !isLetter(s[j]) {
			j--
		}
		if lower(s[i]) != lower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("0P"))
	fmt.Println(string(lower('0')))
}
