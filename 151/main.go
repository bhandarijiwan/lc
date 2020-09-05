// https://leetcode.com/problems/reverse-words-in-a-string/

package main

import (
	"fmt"
	"strings"
)

func reverseWords1(s string) string {
	sb := strings.Builder{}
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	for j := len(arr) - 1; j >= 0; j-- {
		v := false
		for i := 0; i < len(arr[j]); i++ {
			if arr[j][i] != ' ' {
				v = true
				break
			}
		}
		if v {
			// sb.WriteString(strings.TrimSpace(arr[j]))
			sb.WriteString(arr[j])
			if j > 0 {
				sb.WriteByte(' ')
			}
		}
	}
	return sb.String()
}

func rotateLeftN(b []byte, n int) {
	l := len(b)
	j := 0
	for i := n; i < l; i++ {
		b[j] = b[i]
		j += 1
	}
}

func reverse(b []byte) {
	i, j := 0, len(b)-1
	for j > i {
		b[i], b[j] = b[j], b[i]
		j -= 1
		i += 1
	}
}
func Trim(b []byte) []byte {
	i := 0
	for i < len(b) && b[i] == ' ' {
		i++
	}
	j := len(b) - 1
	for j >= 0 && b[j] == ' ' {
		j--
	}
	if i > j {
		return b[0:0]
	}
	b = b[i : j+1]
	i, j = 0, len(b)
	for i < j {
		k := i
		for k < j && b[k] == ' ' {
			k++
		}
		m := k - i
		if m > 1 {
			rotateLeftN(b[i:], m-1)
			j -= m - 1
		}
		i = k - m
		i++
	}
	return b[0:j]
}

func reverseWords(s string) string {
	b := make([]byte, len(s))
	copy(b, s)
	reverse(b)
	b = Trim(b)
	for i := 0; i < len(b); i++ {
		j := i
		for j < len(b) && b[j] != ' ' {
			j++
		}
		reverse(b[i:j])
		i = j
	}
	return string(b)
}

func main() {
	// s := []byte("   a   b  c d   e  ")
	// b := Trim(s)
	// fmt.Println(string(b))
	fmt.Println(reverseWords("hello world"))
	fmt.Println(reverseWords("hello world  hi there    "))
	fmt.Println(reverseWords("  hello world!!  "))
	fmt.Println(reverseWords("  hello    world!!  "))
	fmt.Println(reverseWords("   The sky is blue  *"))
	fmt.Println(reverseWords(" "))
	fmt.Println(reverseWords("   a   b  c d   e  "))
}
