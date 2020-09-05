// https://leetcode.com/problems/design-hit-counter/

package main

import (
	"fmt"
)

type HitCounter struct {
	h [300]int
	l int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{}
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	this.FlushUntil(timestamp)
	this.l = timestamp
	i := (timestamp - 1) % 300
	this.h[i] += 1
}

func (this *HitCounter) FlushUntil(timestamp int) {
	i := (this.l) % 300
	c := timestamp - this.l
	if c > 300 {
		c = 300
	}
	for c > 0 {
		j := i % 300
		this.h[j] = 0
		i++
		c--
	}
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	if timestamp-this.l >= 300 {
		return 0
	}
	r := 0
	// count of items
	c := this.l - clamp(timestamp-300)
	i := clamp(timestamp-1-299) % 300
	for c > 0 {
		j := i % 300
		if this.h[j] > 0 {
			r += this.h[j]
		}
		i++
		c--
	}
	// count till
	return r
}
func clamp(n int) int {
	if n < 0 {
		return 0
	}
	return n
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
func main() {
	obj := Constructor()
	obj.Hit(1)
	obj.Hit(1)
	obj.Hit(1)
	obj.Hit(300)
	fmt.Println(obj.GetHits(300))
}
