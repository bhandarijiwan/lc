// https://leetcode.com/problems/jewels-and-stones/
package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	m := make(map[rune]int, len(S))
	for _, c := range S {
		m[c]++
	}
	count := 0
	for _, c := range J {
		count = count + m[c]
	}
	return count

}
func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))
}
