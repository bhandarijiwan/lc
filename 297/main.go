// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// type Codec struct {
// }

// // Serializes a tree to a single string.
// func (this *Codec) serialize1(root *TreeNode) string {
// 	var sb strings.Builder
// 	q := []*TreeNode{root}
// 	sb.WriteString(getNodeStr(root))
// 	l := len(q)
// 	for l > 0 {
// 		nq, leaf := make([]*TreeNode, l*2), true
// 		for i, j := 0, 0; i < l; i, j = i+1, j+2 {
// 			if q[i] != nil {
// 				nq[j] = q[i].Left
// 				nq[j+1] = q[i].Right
// 				leaf = false
// 			}
// 		}
// 		if leaf {
// 			break
// 		}
// 		q = nq
// 		l = len(q)
// 		for i := 0; i < l; i++ {
// 			sb.WriteString(getNodeStr(q[i]))
// 		}
// 	}
// 	return sb.String()
// }
//
// // Deserializes your encoded data to tree.
// func (this *Codec) deserialize1(data string) *TreeNode {
// 	if len(data) < 1 {
// 		return nil
// 	}
// 	arr := strings.Split(data, ",")
// 	l := len(arr) - 1
// 	nodeArr := make([]*TreeNode, len(arr))
// 	nodeArr[0] = getNode(arr[0])
// 	for j := 1; j < l; j++ {
// 		nodeArr[j] = getNode(arr[j])
// 		i := (j - 1) >> 1
// 		if nodeArr[i] != nil {
// 			lc := (i << 1) + 1
// 			if lc < l {
// 				nodeArr[i].Left = nodeArr[lc]
// 			}
// 			rc := lc + 1
// 			if rc < l {
// 				nodeArr[i].Right = nodeArr[rc]
// 			}
// 		}
// 	}
// 	return nodeArr[0]
// }


type Stack struct {
	data []*TreeNode
	top  int
}

func NewStack() *Stack {
	data := make([]*TreeNode, 1)
	return &Stack{data: data, top: -1}
}

func (s *Stack) Push(item *TreeNode) {
	s.top = s.top + 1
	l := len(s.data)
	if l <= s.top {
		newData := make([]*TreeNode, l*2)
		copy(newData, s.data)
		s.data = newData
	}
	s.data[s.top] = item
}

func (s *Stack) Pop() *TreeNode {
	if s.top >= 0 {
		r := s.data[s.top]
		s.top = s.top - 1
		return r
	}
	panic("Empty stack!!")
}

func (s *Stack) Empty() bool {
	return s.top < 0
}
func (s *Stack) Peek() *TreeNode {
	if s.top >= 0 {
		return s.data[s.top]
	}
	panic("Stack is Emtpy!!")
}


type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func getNodeStr(node *TreeNode) string {
	if node == nil {
		return "null,"
	}
	return strconv.Itoa(node.Val) + ","
}

func getNode(s string) *TreeNode {
	if s == "null" {
		return nil
	}
	v, _ := strconv.Atoi(s)
	return &TreeNode{Val: v}
}

func serializeHelper(node *TreeNode, sb *strings.Builder) {
	if node == nil {
		sb.WriteString(getNodeStr(node))
		return
	}
	serializeHelper(node.Left, sb)
	serializeHelper(node.Right, sb)
	sb.WriteString(getNodeStr(node))
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var sb strings.Builder
	serializeHelper(root, &sb)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strArr := strings.Split(data, ",")
	l := len(strArr) - 1
	stack := NewStack()
	for i := 0; i < l; i++ {
		node := getNode(strArr[i])
		if node == nil {
			stack.Push(node)
			continue
		}
		node.Right = stack.Pop()
		node.Left = stack.Pop()
		stack.Push(node)
	}
	return stack.Pop()
}

func main() {
	// s := "1,2,3,null,null,4,5"
	obj := Constructor()
	t := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 0}}, Right: &TreeNode{Val: 5}}}
	r := obj.serialize(t)
	fmt.Println(r)
	t1 := obj.deserialize(r)
	fmt.Println(t1.Val)
	fmt.Println(t1.Left.Val, t1.Right.Val)
	fmt.Println(t1.Left.Left, t1.Left.Right, t1.Right.Left.Val, t1.Right.Right.Val)
	fmt.Println(t1.Left.Left, t1.Left.Right, t1.Right.Left.Left.Val, t1.Right.Right.Right)

	t = nil
	r = obj.serialize(t)
	fmt.Println(r)
	t1 = obj.deserialize(r)
	fmt.Println(t1)

	t = &TreeNode{Val: 10, Right: &TreeNode{Val: 10, Right: &TreeNode{Val: 10}}}
	r = obj.serialize(t)
	fmt.Println(r)
	t1 = obj.deserialize(r)
	fmt.Println(t1.Val)
	fmt.Println(t1.Left, t.Right.Val)
	fmt.Println(t1.Right.Left, t.Right.Right.Val)

}
