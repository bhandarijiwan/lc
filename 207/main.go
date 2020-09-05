// https://leetcode.com/problems/course-schedule/

package main

import (
	"fmt"
)

type Stack struct {
	data []int
	top  int
}

func (s Stack) IsEmpty() bool {
	return s.top < 0
}

func (s Stack) Peek() int {
	if s.top >= 0 {
		return s.data[s.top]
	}
	panic("Empty stack")
}
func (s *Stack) Pop() int {
	if s.top >= 0 {
		a := s.data[s.top]
		s.top = s.top - 1
		return a
	}
	panic("Empty stack!!")
}
func (s *Stack) Push(c int) {
	s.top = s.top + 1
	if len(s.data) <= s.top {
		newData := make([]int, len(s.data)*2)
		copy(newData, s.data)
		s.data = newData
	}
	s.data[s.top] = c

}

func NewStack() *Stack {
	return &Stack{data: make([]int, 1), top: -1}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjMap := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		u, v := prerequisites[i][0], prerequisites[i][1]
		adjMap[u] = append(adjMap[u], v)
	}
	visited := make(map[int]int)
	for i := 0; i < numCourses; i++ {
		if val, ok := visited[i]; ok && val > 0 {
			continue
		}
		stack := NewStack()
		stack.Push(i)
		for !stack.IsEmpty() {
			u := stack.Peek()
			visited[u] = 1
			j := 0
			for _, v := range adjMap[u] {
				val, ok := visited[v]
				if ok && val == 1 {
					return false
				}
				if !ok && val == 0 {
					stack.Push(v)
					j++
				}
			}
			if j == 0 {
				stack.Pop()
				visited[u] = 2
			}
		}
	}
	return true
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish(3, [][]int{{1, 0}, {0, 2}, {2, 1}}))
	fmt.Println(canFinish(3, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	fmt.Println(canFinish(5, [][]int{{1, 0}, {2, 0}, {2, 1}, {3, 2}, {3, 1}, {4, 3}, {1, 4}}))
	fmt.Println(canFinish(7, [][]int{{1, 0}, {0, 3}, {0, 2}, {3, 2}, {2, 5}, {4, 5}, {5, 6}, {2, 4}}))
	fmt.Println(canFinish(4, [][]int{{0, 1}, {3, 1}, {1, 3}, {3, 2}}))
}
