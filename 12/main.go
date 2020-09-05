// https://leetcode.com/problems/integer-to-roman/

package main

import (
	"fmt"
)

func intToRoman1(num int) string {
	denominators := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	s := ""
	n := num
	for n > 0 {
		for i := 0; i < len(denominators); i++ {
			if n >= denominators[i] {
				s = s + symbols[i]
				n = n - denominators[i]
				break
			}
		}
	}
	return s
}
func intToRoman(num int) string {
	s := make([]byte, 20)
	n := num
	i := 0
	for n > 0 {
		if n >= 1000 {
			s[i] = 77
			n = n - 1000
		} else if n >= 900 {
			s[i] = 67
			i = i + 1
			s[i] = 77
			n = n - 900
		} else if n >= 500 {
			s[i] = 68
			n = n - 500
		} else if n >= 400 {
			s[i] = 67
			i = i + 1
			s[i] = 68
			n = n - 400
		} else if n >= 100 {
			s[i] = 67
			n = n - 100
		} else if n >= 90 {
			s[i] = 88
			i = i + 1
			s[i] = 67
			n = n - 90
		} else if n >= 50 {
			s[i] = 76
			n = n - 50
		} else if n >= 40 {
			s[i] = 88
			i = i + 1
			s[i] = 76
			n = n - 40
		} else if n >= 10 {
			s[i] = 88
			n = n - 10
		} else if n >= 9 {
			s[i] = 73
			i = i + 1
			s[i] = 88
			n = n - 9
		} else if n >= 5 {
			s[i] = 86
			n = n - 5
		} else if n >= 4 {
			s[i] = 73
			i = i + 1
			s[i] = 86
			n = n - 4
		} else {
			s[i] = 73
			n = n - 1
		}
		i = i + 1
	}
	return string(s[0:i])
}

func main() {
	fmt.Println(intToRoman(3999))
}
