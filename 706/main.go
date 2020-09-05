// https://leetcode.com/problems/design-hashmap/

package main

import (
	"fmt"
)

const CAPACITY = 100

type Element struct {
	key int
	val int
	del bool
}

type MyHashMap struct {
	m [100][]Element
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	bucketIndex := key % CAPACITY
	bucket := this.m[bucketIndex]
	l := len(bucket)
	d := l
	for j := 0; j < l; j++ {
		if bucket[j].del && d == l {
			d = j
		}
		if bucket[j].key == key {
			d = j
			break
		}
	}
	if d < l {
		bucket[d].key = key
		bucket[d].val = value
		bucket[d].del = false
	} else {
		bucket = append(bucket, Element{key: key, val: value, del: false})
	}
	this.m[bucketIndex] = bucket
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	bucketIndex := key % CAPACITY
	bucket := this.m[bucketIndex]
	l := len(bucket)
	for i := 0; i < l; i++ {
		if !bucket[i].del && bucket[i].key == key {
			return bucket[i].val
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	bucketIndex := key % CAPACITY
	bucket := this.m[bucketIndex]
	l := len(bucket)
	for i := 0; i < l; i++ {
		if bucket[i].key == key {
			bucket[i].del = true
		}
	}
}

func main() {
	h := Constructor()
	fmt.Println(h.Get(10))
	h.Put(1000000, 20)
	fmt.Println(h.Get(10))
	h.Put(10, 30)
	fmt.Println(h.Get(10))
	h.Remove(10)
	fmt.Println(h.Get(1000000))
	fmt.Println(h.Get(2000000))
	fmt.Println(h.m)
}
