package main

import (
	"fmt"
	"sync"
)
var sw sync.WaitGroup
func T1() {
	fmt.Println("Hello from T1")
}

func T2() {
	fmt.Println("Hello from T2")
}

func main() {
	sw.Add(2)
	go T1()
	go T2()
	sw.Wait()
}
