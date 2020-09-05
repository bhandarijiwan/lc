// https://leetcode.com/problems/logger-rate-limiter/

package main

import (
	"fmt"
)

type Logger struct {
	h map[string]int
}

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{h: make(map[string]int)}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	last := this.h[message]
	if last >= 1 && timestamp-(last-1) < 10 {
		return false
	}
	this.h[message] = timestamp + 1
	return true
}

func main() {
	obj := Constructor()
	fmt.Println(obj.ShouldPrintMessage(1, "foo"))
	fmt.Println(obj.ShouldPrintMessage(11, "foo"))
}
