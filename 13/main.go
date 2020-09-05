//https://leetcode.com/problems/roman-to-integer/
package main

import (
	"fmt"
)

func romanToInt(s string) int {
	l := len(s)
	symbol := make(map[string]int, l)
	symbol["I"] = 1
	symbol["V"] = 5
	symbol["X"] = 10
	symbol["L"] = 50
	symbol["C"] = 100
	symbol["D"] = 500
	symbol["M"] = 1000
	symbol["IV"] = 4
	symbol["IX"] = 9
	symbol["XL"] = 40
	symbol["XC"] = 90
	symbol["CD"] = 400
	symbol["CM"] = 900

	v, i := 0, 0
	for i < l {
		if i+1 < l {
			if n, ok := symbol[s[i:i+2]]; ok {
				v = v + n
				i = i + 1
			} else {
				v = v + symbol[s[i:i+1]]
			}
		} else {
			v = v + symbol[s[i:i+1]]
		}
		i = i + 1
	}
	return v
}

func main() {
	fmt.Println(romanToInt("MCDLXXVI"))
}
