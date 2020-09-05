//https://leetcode.com/problems/reverse-string/

package main

import "fmt"

func reverseString(s []byte) {
	l := len(s) - 1
	i := 0
	for l > i {
		a := s[i]
		s[i] = s[l]
		s[l] = a
		i++
		l--
	}

}

func main() {
	input := []byte("hello")
	reverseString(input)
	fmt.Printf("%s\n", input)
	input = []byte("Hadn")
	reverseString(input)
	fmt.Printf("%s\n", input)
}
