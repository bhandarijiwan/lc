// https://leetcode.com/problems/regular-expression-matching/
package main

import (
	"fmt"
)


func isMatch1(s string, p string) bool {
	lenS, lenP := len(s), len(p)
	if lenP < 1 {
		return lenS < 1
	}
	m := lenS > 0 && (s[0] == p[0] || p[0] == '.')
	if lenP >= 2 && p[1] == '*' {
		return isMatch1(s, p[2:]) || (m && isMatch1(s[1:], p))
	} else {
		return m && isMatch1(s[1:], p[1:])
	}
}
func match(a, b byte) bool {
	if b == '.' {
		return true
	}
	return a == b
}

func isMatch(s string, p string) bool {
	lenS, lenP := len(s), len(p)
	dp := make([][]bool, lenS+1)
	for i := 0; i < lenS+1; i++ {
		dp[i] = make([]bool, lenP+1)
	}
	dp[0][0] = true
	// pattern of a*, a*b* etc match empty string so
	for i := 1; i <= lenP; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}
	for i := 1; i <= lenS; i++ {
		for j := 1; j <= lenP; j++ {
			if match(s[i-1], p[j-1]) {
				dp[i][j] = dp[i-1][j-1]
			} else {
				if p[j-1] == '*' {
					dp[i][j] = (match(s[i-1], p[j-2]) && dp[i-1][j]) || dp[i][j-2]
				}
			}
		}
	}
	return dp[lenS][lenP]
}

func assert(a bool) {
	if !a {
		panic("Failed")
	}
	fmt.Println("Looks good!!")
}

func main() {
	assert(isMatch("aa", "a") == false)
	assert(isMatch("aa", "a*") == true)
	assert(isMatch("aa", ".*") == true)
	assert(isMatch("aab", "c*a*b") == true)
	assert(isMatch("ippi", "p*.") == false)
	assert(isMatch("mississippi", "mis*is*p*.") == false)
	assert(isMatch("aaa", "a*aaa") == true)
}
