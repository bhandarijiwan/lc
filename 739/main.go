// https://leetcode.com/problems/daily-temperatures/

package main

import (
	"fmt"
)

func dailyTemperatures(T []int) []int {
	l := len(T)
	r := make([]int, l)
	for j := l - 2; j >= 0; j-- {
		if T[j+1] > T[j] {
			r[j] = 1
		} else {
			if r[j+1] > 0 {
				i := j + r[j+1] + 1
				for i < l {
					if T[i] > T[j] {
						r[j] = i - j
						break
					}
					if r[i] < 1 {
						break
					}
					i = i + r[i]
				}
			}
		}
	}
	return r
}

func main() {
	fmt.Println(dailyTemperatures([]int{78, 62, 50, 71, 78, 82}))
}
