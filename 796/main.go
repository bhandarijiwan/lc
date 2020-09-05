// https://leetcode.com/problems/rotate-string/

package main

import (
	"fmt"
)

func rotate(str string) string {
	l := len(str)
	r := make([]byte, l)
	for i := 1; i <= len(str); i++ {
		r[(i-1)%l] = str[i%l]
	}
	return string(r)
}

func rotateString1(A string, B string) bool {
	a := A
	for i := 0; i < len(A); i++ {
		a = rotate(a)
		if a == B {
			return true
		}
	}
	if len(A) < 1 && len(B) < 1 {
		return true
	}
	return false
}

func rotateString2(A string, B string) bool {
	lA, lB := len(A), len(B)
	if lA == 0 && lB == 0 {
		return true
	}
	if lA != lB {
		return false
	}
	match := false
	i, j := 0, 0
	for match || j < lB {
		if B[j%lB] == A[i] {
			i++
			match = true
		} else {
			match = false
			j = j - i
			i = 0
		}
		if i >= lA {
			return true
		}
		j = j + 1
	}
	return false
}

func rotateString(A string, B string) bool {
	// build a dfa/prefix table on the pattern
	
}

func main() {
	fmt.Println(rotateString("abcde", "cdeab"))
	fmt.Println(rotateString("abcde", "abced"))
	fmt.Println(rotateString("", ""))
	fmt.Println(rotateString("bbbacddceeb", "ceebbbbacdd"))

}
