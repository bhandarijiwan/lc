// https://leetcode.com/problems/coin-change/

package main

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	l := len(coins)
	if l < 1 {
		return -1
	}
	if amount < 1 {
		return amount
	}
	x := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		x[i] = -1
		if i >= coins[0] && i%coins[0] == 0 {
			x[i] = i / coins[0]
		}
	}
	for i := 1; i < l; i++ {
		for j := 0; j <= amount; j++ {
			if j > coins[i] {
				p := x[j]
				c := x[j-coins[i]] + 1
				if p > 0 && c > 0 {
					if c < p {
						x[j] = c
					}
				} else if p > 0 || c > 0 {
					if c > p {
						x[j] = c
					}
				}
			} else if j == coins[i] {
				x[j] = 1
			}
		}
	}
	return x[amount]
}

func main() {
	fmt.Println(coinChange([]int{}, 1))
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1, 5, 10}, 50))
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
	fmt.Println(coinChange([]int{5, 4, 2}, 8))
}
