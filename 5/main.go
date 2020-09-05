// https://leetcode.com/problems/longest-palindromic-substring/

package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}
	// t := l / 2
	u, v, m := 0, 1, 0
	for i := 0; i < l; i++ {
		//
		// fmt.Println("For char ", string(s[i]))
		// fmt.Println("Center 1")
		j, k := i-1, i+1
		for j >= 0 && k < l && s[j] == s[k] {
			j -= 1
			k += 1
		}
		c := k - (j + 1)
		if c > m {
			m = c
			u, v = j+1, k
		}
		// fmt.Println(s[j+1 : k])
		// fmt.Println("length=", k-(j+1))
		// fmt.Println("Center 2")
		j, k = i, i+1
		for j >= 0 && k < l && s[j] == s[k] {
			j -= 1
			k += 1
		}
		c = k - (j + 1)
		if c > m {
			m = c
			u, v = j+1, k
		}
		// fmt.Println("length=", k-(j+1))
	}

	return s[u:v]
}

func main() {
	fmt.Println(longestPalindrome("babadd"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome(""))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("aabbaa"))
	fmt.Println(longestPalindrome("eabcb"))
}
