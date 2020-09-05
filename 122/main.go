// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	l := len(prices)
	md := 0
	p := 0
	for j := l - 2; j >= 0; j-- {
		k := prices[j+1] - prices[j]
		if k > 0 {
			p = md + k
			if p > md {
				md = p
			}
		}
	}
	return p
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
