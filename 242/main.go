//https://leetcode.com/problems/valid-anagram/

package main

import "fmt"

func isAnagram(s string, t string) bool {
	m := make([]int, 26)
	for _, c := range s {
		m[c-'a']++
	}
	for _, c := range t {
		m[c-'a']--
	}
	for _, c := range m {
		if c != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
