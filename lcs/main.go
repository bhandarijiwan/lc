// longest common subsequence
package main

import (
	"fmt"
)

func lcsI(strA string, strB string, i, j int, s [][]string) string {
	if i < 0 || j < 0 {
		return ""
	}
	if s[i][j] != "#" {
		return s[i][j]
	}
	str := ""
	if strA[i] == strB[j] {
		str = lcsI(strA, strB, i-1, j-1, s) + string(strA[i])
	} else {
		s1 := lcsI(strA, strB, i-1, j, s)
		s2 := lcsI(strA, strB, i, j-1, s)
		if len(s1) > len(s2) {
			str = s1
		} else {
			str = s2
		}
	}
	s[i][j] = str
	return str
}

func lcs(strA string, strB string) string {
	s := make([][]string, len(strA))
	for i := 0; i < len(strA); i++ {
		s[i] = make([]string, len(strB))
		for j := 0; j < len(strB); j++ {
			s[i][j] = "#"
		}
	}
	return lcsI(strA, strB, len(strA)-1, len(strB)-1, s)
}

func main() {
	fmt.Println(lcs("AGBHCIDJEK", "AFOPBCDE"))
}
