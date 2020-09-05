// https://leetcode.com/problems/my-calendar-i/

package main

import (
	"fmt"
	"sort"
)

type MyCalendar struct {
	events [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{events: make([][]int, 0)}
}

func overlaps(a, b []int) bool {
	if b[0] >= a[1] || a[0] >= b[1] {
		return false
	}
	return true
}

func (this *MyCalendar) exists(start, end int) (bool, int) {
	b := []int{start, end}
	index := sort.Search(len(this.events), func(i int) bool {
		if overlaps(this.events[i], b) {
			return true
		}
		return this.events[i][0] > b[0]
	})
	if index < len(this.events) && overlaps(this.events[index], b) {
		return true, index
	}
	return false, index
}

func (this *MyCalendar) Book(start int, end int) bool {
	ok, index := this.exists(start, end)
	if !ok {
		this.events = append(this.events, []int{start, end})
		l := len(this.events)
		this.events[l-1], this.events[index] = this.events[index], this.events[l-1]
		index = index + 1
		for index < l {
			this.events[l-1], this.events[index] = this.events[index], this.events[l-1]
			index = index + 1
		}
	}
	return !ok
}

func main() {
	cal := Constructor()
	fmt.Println(cal.Book(0, 1))
	fmt.Println(cal.Book(10, 20))
	fmt.Println(cal.Book(10, 20))
	fmt.Println(cal.Book(1, 6))
	fmt.Println(cal.Book(6, 7))
	fmt.Println(cal.events)
}
