//https://leetcode.com/problems/palindrome-number/

package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	count := 0
	n := x
	for n > 0 {
		n = n / 10
		count++
	}
	d := 1
	for i := 0; i < count-1; i++ {
		d = d * 10
	}
	n = x
	for d > 0 && n >= 10 {
		l := n / d
		n = n % d
		r := n % 10
		if l != r {
			return false
		}
		n = n / 10
		d = d / 100
	}
	if d >= 10 && n >= 1 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(989))
}
