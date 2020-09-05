//https://leetcode.com/problems/most-common-word/

package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

func toLowerAndRemovePunct(s string) string {
	m := make([]byte, len(s))
	j := 0
	for _, c := range s {
		if !unicode.IsPunct(c) {
			m[j] = byte(c)
			if m[j] >= 'A' && m[j] <= 'Z' {
				m[j] = m[j] + ('a' - 'A')
			}
		} else {
			m[j] = ' '
		}
		j++
	}
	return string(m)
}

func mostCommonWord(paragraph string, banned []string) string {
	m := make(map[string]bool)
	for _, s := range banned {
		m[s] = true
	}
	s := bufio.NewScanner(strings.NewReader(toLowerAndRemovePunct(paragraph)))
	s.Split(bufio.ScanWords)
	wordsMap := make(map[string]int)
	r := ""
	maxCount := 0
	for s.Scan() {
		str := s.Text()
		str = strings.ToLower(str)
		if _, ok := m[str]; !ok {
			wordsMap[str]++
			if wordsMap[str] > maxCount {
				r = str
				maxCount = wordsMap[str]
			}
		}
	}
	return r
}

func main() {

	p := "Bob,boo!"
	banned := []string{"hit!"}
	fmt.Println(mostCommonWord(p, banned))

}
