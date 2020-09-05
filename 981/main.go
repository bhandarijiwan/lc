// https://leetcode.com/problems/time-based-key-value-store/

package main

import (
	"fmt"
)

type Data struct {
	timestamp int
	val       string
}

type TimeMap struct {
	v map[string][]Data
}

func (t *TimeMap) search(key string, n int) int {
	src, ok := t.v[key]
	if !ok {
		return -1
	}
	i, j := 0, len(src)-1
	for j >= i {
		m := i + (j-i)/2
		if src[m].timestamp == n {
			return m
		} else if n > src[m].timestamp {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	if i < j {
		return i
	}
	return j
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{v: make(map[string][]Data)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	d := Data{timestamp: timestamp, val: value}
	this.v[key] = append(this.v[key], d)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	index := this.search(key, timestamp)
	if index < 0 {
		return ""
	}
	return this.v[key][index].val
}

func main() {
	obj := Constructor()
	obj.Set("love", "high", 10)
	obj.Set("love", "low", 20)
	fmt.Println(obj.Get("love", 5))
	fmt.Println(obj.Get("love", 10))
	fmt.Println(obj.Get("love", 15))
	fmt.Println(obj.Get("love", 20))
	fmt.Println(obj.Get("love", 25))
}
