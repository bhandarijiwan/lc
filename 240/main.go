// https://leetcode.com/problems/search-a-2d-matrix-ii/
// # binary search, #search space reduction, # bottom right
package main

import (
	"fmt"
	"sort"
)

func searchMatrix1(matrix [][]int, target int) bool {

	// binary search
	for i := 0; i < len(matrix); i++ {
		v := matrix[i]
		l := len(v)
		if l > 0 && v[0] > target {
			break
		}
		i := sort.Search(l, func(i int) bool {
			return v[i] >= target
		})
		if i < l && v[i] == target {
			return true
		}
	}
	return false
}

func binarySearch(nums [][]int, target int, row bool) (bool, int) {
	i, j := 0, len(nums)-1
	if !row {
		j = len(nums[0]) - 1
	}
	for j >= i {
		m := i + (j-i)/2
		r, c := m, 0
		if !row {
			r, c = 0, m
		}
		if nums[r][c] == target {
			return true, m
		} else if nums[r][c] > target {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	if i > j {
		return false, j
	}
	return false, i
}

func searchMatrix2(matrix [][]int, target int) bool {

	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	ok, rowIndex := binarySearch(matrix, target, true)
	if ok {
		return true
	}
	if rowIndex < 0 {
		return false
	}
	index := sort.Search(len(matrix[rowIndex]), func(i int) bool {
		return matrix[rowIndex][i] >= target
	})
	if index < len(matrix[rowIndex]) && matrix[rowIndex][index] == target {
		return true
	}
	ok, colIndex := binarySearch(matrix, target, false)
	if ok {
		return true
	}
	index = sort.Search(len(matrix), func(i int) bool {
		return matrix[i][colIndex] >= target
	})
	if index < len(matrix) && matrix[index][colIndex] == target {
		return true
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	r, c := m-1, 0
	for r >= 0 && c < n {
		if matrix[r][c] == target {
			return true
		} else if target > matrix[r][c] {
			c++
		} else {
			r--
		}
	}
	return false
}

func main() {
	input := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(searchMatrix(input, 5))
	fmt.Println(searchMatrix(input, 8))
	fmt.Println(searchMatrix(input, 6))
	fmt.Println(searchMatrix(input, 1))
	fmt.Println(searchMatrix(input, 3))
	fmt.Println(searchMatrix(input, 4))
	fmt.Println(searchMatrix(input, 12))
	fmt.Println(searchMatrix(input, 112))
	fmt.Println(searchMatrix([][]int{{12}}, 1))
	fmt.Println(searchMatrix([][]int{{}}, 1))
	input = [][]int{{1, 3, 5, 7, 9}, {2, 4, 6, 8, 10}, {11, 13, 15, 17, 19}, {12, 14, 16, 18, 20}, {21, 22, 23, 24, 25}}
	fmt.Println(searchMatrix(input, 13))
}
