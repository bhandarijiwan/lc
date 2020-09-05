// https://leetcode.com/problems/reorder-data-in-log-files/

package main

import (
	"fmt"
	"reflect"
	"sort"
)

func getIndex(str string) int {
	j := 0
	for ; str[j] == ' '; j++ {
	}
	for ; str[j] != ' '; j++ {
	}
	for ; str[j] == ' '; j++ {
	}
	return j
}
func reorderLogFiles(logs []string) []string {
	j := len(logs) - 1
	for i := j; i >= 0; i-- {
		m := getIndex(logs[i])
		s1 := logs[i][m:]
		if isDigitLog(s1) {
			logs[i], logs[j] = logs[j], logs[i]
			j--
		}
	}
	fmt.Println(logs[:j+1])
	sort.Slice(logs[0:j], func(i, j int) bool {
		m, n := getIndex(logs[i]), getIndex(logs[j])
		s1, s2 := logs[i][m:], logs[j][n:]
		if s1 == s2 {
			return logs[i][:m] < logs[j][:n]
		}
		return s1 < s2
	})
	return logs
}

func reorderLogFiles1(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		m, n := getIndex(logs[i]), getIndex(logs[j])
		s1, s2 := logs[i][m:], logs[j][n:]
		if !isDigitLog(s1) && !isDigitLog(s2) {
			if s1 == s2 {
				return logs[i][:m] < logs[j][:n]
			}
			return s1 < s2
		}
		if isDigitLog(s1) {
			if isDigitLog(s2) {
				return false
			}
			return false
		}
		return true
	})
	fmt.Println(logs)
	return logs
}

func isDigitLog(s string) bool {
	return s[0] >= '0' && s[0] <= '9'
}

func main() {
	logs := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
	// logs = []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo", "a2 act car"}
	// logs = []string{"mt awiizp ubo w", "rs 57969 673 9", "94 zhrucvisya h", "z 3445409 12460", "wj5 5907 615745", "l 4135283 38 90", "sei mybsw gf d", "ja kwygowpegz w", "0z2 1 25410 8 9", "0qj ksxe fb nbl", "21 nadrjznr m v", "o mhvxjq psn tv", "h 8597584 44 07", "qch uaw eezjvaz", "v 38343357670 3"}
	// expect := []string{"mt awiizp ubo w", "0qj ksxe fb nbl", "ja kwygowpegz w", "o mhvxjq psn tv", "sei mybsw gf d", "21 nadrjznr m v", "qch uaw eezjvaz", "94 zhrucvisya h", "rs 57969 673 9", "z 3445409 12460", "wj5 5907 615745", "l 4135283 38 90", "0z2 1 25410 8 9", "h 8597584 44 07", "v 38343357670 3"}
	logs = []string{"8 fj dzz k", "5r 446 6 3", "63 gu psub", "5 ba iedrz", "6s 87979 5", "3r 587 01", "jc 3480612", "bb wsrd kp", "b aq cojj", "r5 6316 71"}
	actual := reorderLogFiles1(logs)
	fmt.Println(reflect.DeepEqual(logs, actual))

}
