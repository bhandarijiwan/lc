//https://leetcode.com/problems/palindrome-permutation/
package main

import "fmt"

func canPermutePalindrome(s string) bool {
	m := make(map[rune]int)
	oddCount := 0
	for _, r := range s {
		k := m[r]
		m[r] = k + 1
		if (k+1)%2 == 0 {
			oddCount--
		} else {
			oddCount++
		}
	}
	if oddCount > 1 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canPermutePalindrome("aab"))
}
