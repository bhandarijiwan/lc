// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	n *TreeNode
	p *Node
	l int
}

func makeTree(root *TreeNode, parent *Node, level int, np, nq **Node) {
	if root == nil {
		return
	}
	node := &Node{n: root, p: parent, l: level}
	p, q := (*np).n, (*nq).n
	if root == p {
		*np = node
	}
	if root == q {
		*nq = node
	}
	if (*np).p != nil && (*nq).p != nil {
		return
	}
	makeTree(root.Left, node, level+1, np, nq)
	makeTree(root.Right, node, level+1, np, nq)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	np, nq := &Node{n: p}, &Node{n: q}
	makeTree(root, nil, 0, &np, &nq)
	low, high := np, nq
	if np.l < nq.l {
		low, high = nq, np
	}
	for low.l > high.l {
		low = low.p
	}
	for low != high {
		low = low.p
		high = high.p
	}
	return low.n
}

const (
	WHITE = 0
	BLACK = 1
)

type DSNode struct {
	*TreeNode
	p *DSNode
	a *DSNode
	c int8
	r int
}

// MakeSet makes a new node
func MakeSet(u *TreeNode) *DSNode {
	n := &DSNode{}
	n.TreeNode = u
	n.p = n
	return n
}

// FindSet the parent of a node and also does path compression
func FindSet(node *DSNode) *DSNode {
	if node.p == node {
		return node
	}
	node.p = FindSet(node.p)
	return node.p
}

// Union merges the set represented by FindSet(U) and the set
// represented by FindSet(V)
func Union(u, v *DSNode) {
	x, y := FindSet(u), FindSet(v)
	if x.r > y.r {
		y.p = x
	} else {
		x.p = y
		if x.r == y.r {
			y.r += 1
		}
	}
}

func lca1(r, p, q *TreeNode) *DSNode {
	if r == nil {
		return nil
	}
	u := MakeSet(r)
	FindSet(u).a = u
	if r.Left != nil {
		left := lca1(r.Left, p, q)
		Union(u, left)
		FindSet(u).a = u
	}
	if r.Right != nil {
		right := lca1(r.Right, p, q)
		Union(u, right)
		FindSet(u).a = u
	}
	u.c = BLACK
	return u
}

func lca(r, p, q *TreeNode) *TreeNode {
	// convert r to an array
	// +-1 Range Minimum Query

	// let's create an array of depths
	//
}

func main() {
	m := &TreeNode{Val: 8}
	n := &TreeNode{Val: 7}
	o := &TreeNode{Val: 4}
	p := &TreeNode{Val: 5,
		Left: &TreeNode{Val: 6},
		Right: &TreeNode{Val: 2,
			Left:  n,
			Right: o,
		},
	}
	q := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 0},
		Right: m,
	}
	t := &TreeNode{Val: 3,
		Left:  p,
		Right: q,
	}

	// fmt.Println(lowestCommonAncestor(t, n, q).Val)
	r := lca1(t, n, q)
	fmt.Println(r.c)
	fmt.Println(r.p.Val)
}
