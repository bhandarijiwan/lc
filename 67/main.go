// https://leetcode.com/problems/add-binary/

package main

import (
	"fmt"
)

func addBit(a, b byte) (byte, byte) {
	r := a ^ b + 48
	if a == '1' && b == '1' {
		return r, '1'
	}
	return r, '0'
}

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	l := lb
	if lb < la {
		l = la
	}
	c, c1, c2 := byte('0'), byte('0'), byte('0')
	r := make([]byte, l+1)
	k := len(r) - 1
	var i, j int
	for i, j = la-1, lb-1; i >= 0 && j >= 0; {
		r[k], c1 = addBit(a[i], c)
		r[k], c2 = addBit(r[k], b[j])
		c = c1 ^ c2 + 48
		i--
		j--
		k--
	}
	if j >= 0 {
		for j >= 0 {
			r[k], c = addBit(b[j], c)
			j--
			k--
		}
	}
	if i >= 0 {
		for i >= 0 {
			r[k], c = addBit(a[i], c)
			i--
			k--
		}
	}
	if c == '1' {
		r[k] = c
	} else {
		k++
	}
	return string(r[k:])
}

func main() {
	r := addBinary("1111", "1111")
	fmt.Println(r)
}
