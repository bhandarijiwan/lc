//https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

package main

import "fmt"

func maxProfit(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	prev := 0
	md := prices[l-1]
	for i := l - 2; i >= 0; i-- {
		d := md - prices[i]
		if d > prev {
			prev = d
		}
		if prices[i] > md {
			md = prices[i]
		}
	}
	return prev
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{1, 2}))
	fmt.Println(maxProfit([]int{1}))
	fmt.Println(maxProfit([]int{}))
}
