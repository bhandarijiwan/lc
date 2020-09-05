//https://leetcode.com/problems/happy-number/

package main

import "fmt"

func getDigitsSquaredSum(n int) int {
	s := 0
	for n > 0 {
		d := n % 10
		n = n / 10
		s = s + d*d
	}
	return s
}

func isHappy(n int) bool {
	h := make(map[int]bool)
	for !h[n] {
		h[n] = true
		if n = getDigitsSquaredSum(n); n == 1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(78))
}
