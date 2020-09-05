//https://leetcode.com/problems/first-unique-character-in-a-string/

package main

import "fmt"

func firstUniqChar(s string) int {
	m := make([]rune, 26)
	for _, c := range s {
		m[c-'a']++
	}
	for i, c := range s {
		if m[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}
