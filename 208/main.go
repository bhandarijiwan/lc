// https://leetcode.com/problems/implement-trie-prefix-tree/

package main

import (
	"fmt"
)

type Node struct {
	children [26]*Node
	final    bool
}

type Trie struct {
	root *Node
}

func NewNode() *Node {
	return &Node{}
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: NewNode(),
	}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		index := char - 'a'
		nextNode := node.children[index]
		if nextNode == nil {
			nextNode = NewNode()
			node.children[index] = nextNode
		}
		node = nextNode
	}
	node.final = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		index := char - 'a'
		nextNode := node.children[index]
		if nextNode == nil {
			return false
		}
		node = nextNode
	}
	return node.final == true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, char := range prefix {
		index := char - 'a'
		node = node.children[index]
		if node == nil {
			return false
		}
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("hello")
	fmt.Println(trie.Search("hell"))
	fmt.Println(trie.Search("helloa"))
	trie.Insert("hellloa")
	fmt.Println(trie.StartsWith("hellloa"))
	fmt.Println(trie.Search("hellloa"))
}