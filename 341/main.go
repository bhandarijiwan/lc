// https://leetcode.com/problems/flatten-nested-list-iterator/

package main



// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {}

type NestedIterator struct {
	list  []*NestedInteger
	index int
}

func flatten(nestedList []*NestedInteger, flatlist *[]*NestedInteger) {
	for _, ni := range nestedList {
		if ni.IsInteger() {
			*flatlist = append(*flatlist, ni)
		} else {
			flatten(ni.GetList(), flatlist)
		}
	}
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	flatlist := make([]*NestedInteger, 0)
	flatten(nestedList, &flatlist)
	return &NestedIterator{list: flatlist, index: 0}
}

func (this *NestedIterator) Next() int {
	l := len(this.list)
	if this.index >= l {
		return this.list[l-1].GetInteger()
	}
	item := this.list[this.index]
	this.index = this.index + 1
	return item.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	l := len(this.list)
	if this.index >= l {
		return false
	}
	return true
}

func main() {

}
