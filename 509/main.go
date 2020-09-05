//https://leetcode.com/problems/fibonacci-number/

package main

import "fmt"

func fib(N int) int {
	x := 0
	y := 1
	if N < 2 {
		if N > 0 {
			return y
		}
		return  x
	}
	i := 0
	for i < N {
		a := y
		y = y + x
		x = a
		i++
	}
	return x
}

func main() {
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
}
