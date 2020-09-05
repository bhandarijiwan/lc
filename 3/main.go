// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	i, j, l := 0, 0, len(s)
	r := 0
	for j = 0; j < l; j++ {
		if idx, ok := m[s[j]]; ok && idx >= i {
			if j-i > r {
				r = j - i
			}
			i = m[s[j]] + 1
		}
		m[s[j]] = j
	}
	if j-i > r {
		r = j - i
	}
	return r
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb") == 3)
	fmt.Println(lengthOfLongestSubstring("bbbbbb") == 1)
	fmt.Println(lengthOfLongestSubstring("pwwkew") == 3)
	fmt.Println(lengthOfLongestSubstring("pwwkews") == 4)
	fmt.Println(lengthOfLongestSubstring("pwwk") == 2)
	fmt.Println(lengthOfLongestSubstring("pwwk") == 2)
	fmt.Println(lengthOfLongestSubstring("ababbac") == 3)
	fmt.Println(lengthOfLongestSubstring("a") == 1)
	fmt.Println(lengthOfLongestSubstring("abc") == 3)
	fmt.Println(lengthOfLongestSubstring("babbbcd") == 3)
	fmt.Println(lengthOfLongestSubstring("") == 0)
}
