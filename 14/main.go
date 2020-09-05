// https://leetcode.com/problems/longest-common-prefix/

package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	l := 2<<32 - 1
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < l {
			l = len(strs[i])
		}
	}
	i := 0
outer:
	for i = 0; i < l; i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != c {
				break outer
			}
		}
	}
	return strs[0][0:i]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"", ""}))
}
