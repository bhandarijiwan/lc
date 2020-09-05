// https://leetcode.com/problems/utf-8-validation/
package main

import (
	"fmt"
)

const c1, c2 = 0x80, 0xBF
const l0, h0 = 0x0, 0x7F  // start of 1 byte sequence
const l1, h1 = 0xC0, 0xDF // start of 2 byte sequence
const l2, h2 = 0xE0, 0xEF // start of 3 byte sequence
const l3, h3 = 0xF0, 0xF7 // start of 4 byte sequence

func checkContinuation(b int) bool {
	if b >= c1 && b <= c2 {
		return true
	}
	return false
}

func getSize(b int) (bool, int) {
	if b >= l0 && b <= h0 {
		return true, 0
	} else if b >= l1 && b <= h1 {
		return true, 1
	} else if b >= l2 && b <= h2 {
		return true, 2
	} else if b >= l3 && b <= h3 {
		return true, 3
	}
	return false, 0
}

func validUtf8(data []int) bool {
	l := len(data)
	size, c := 0, 0
	for i := 0; i < l; i++ {
		switch size {
		case 1, 2, 3:
			if !checkContinuation(data[i]) {
				return false
			}
			c += 1
			break
		default:
			if size != 0 {
				return false
			}
			ok, s := getSize(data[i])
			if !ok || i+s >= l {
				return false
			}
			size = s
		}
		if c > size {
			return false
		} else if c == size {
			c = 0
			size = 0
		}
	}
	return true
}

func main() {
	fmt.Println(validUtf8([]int{197, 130, 1}))
	fmt.Println(validUtf8([]int{235, 140, 4}))
	fmt.Println(validUtf8([]int{235}))
}
