// https://leetcode.com/problems/find-all-anagrams-in-a-string/

package main

import (
	"fmt"
)

func checkAnagram(a []int16, b []int16) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	ls, lp := len(s), len(p)
	if lp > ls {
		return []int{}
	}
	k := make([]int16, 26)
	for i := 0; i < lp; i++ {
		k[p[i]-'a'] += 1
	}
	w := make([]int16, 26)
	r := make([]int, 0)
	for i := 0; i < lp; i++ {
		w[s[i]-'a'] += 1
	}
	// check anagram
	if checkAnagram(k, w) {
		r = append(r, 0)
	}
	for i := lp; i < ls; i++ {
		w[s[i-lp]-'a'] -= 1
		w[s[i]-'a'] += 1
		if checkAnagram(k, w) {
			r = append(r, i-lp+1)
		}
	}
	return r
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
	fmt.Println(findAnagrams("", "a"))
	fmt.Println(findAnagrams("", ""))
}
