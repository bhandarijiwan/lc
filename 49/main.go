// https://leetcode.com/problems/group-anagrams/

package main

import (
	"bytes"
	"fmt"
)

func groupAnagrams1(strs []string) [][]string {
	r := make([][]string, 0)
	i := 0
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	m := make(map[int]int)
	for _, str := range strs {
		key := 1
		for _, c := range str {
			key = key * primes[int(c-'a')]
		}
		if index, ok := m[key]; ok {
			r[index] = append(r[index], str)
		} else {
			r = append(r, []string{str})
			m[key] = i
			i = i + 1
		}
	}
	return r
}

func transform(str string) string {
	b := bytes.Repeat([]byte{'0'}, 26)
	for i := 0; i < len(str); i++ {
		v := b[str[i]-'a']
		v = v + 1
		b[str[i]-'a'] = v
	}
	return string(b)
}

func groupAnagrams(strs []string) [][]string {
	r := make([][]string, 0)
	i := 0
	m := make(map[string]int)
	for _, str := range strs {
		key := transform(str)
		if index, ok := m[key]; ok {
			r[index] = append(r[index], str)
		} else {
			r = append(r, []string{str})
			m[key] = i
			i = i + 1
		}
	}
	return r
}

func main() {
	input := []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}
	fmt.Println(groupAnagrams(input))
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
