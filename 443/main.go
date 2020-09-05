//https://leetcode.com/problems/string-compression/
package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	l := len(chars)
	if l < 1 {
		return l
	}
	p, c, j := chars[0], 0, 0
	for i := 0; i < l; i++ {
		if chars[i] != p {
			chars[j] = p
			if c > 1 {
				j = j + 1
				b := []byte(strconv.Itoa(c))
				copy(chars[j:], b)
				j = j + len(b) - 1
			}
			j = j + 1
			c = 1
		} else {
			c = c + 1
		}
		p = chars[i]
	}
	chars[j] = p
	if c > 1 {
		j = j + 1
		b := []byte(strconv.Itoa(c))
		copy(chars[j:], b)
		j = j + len(b) - 1
	}
	return j + 1
}

func main() {
	input := []byte("aaadaaabaa")
	r := compress(input)
	fmt.Println(r)
	fmt.Println(string(input[:r]))
}
