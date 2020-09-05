// https://leetcode.com/problems/sparse-matrix-multiplication/

package main

import (
	"fmt"
)

func transpose(m [][]int) [][]int {
	r := len(m)
	if r < 1 {
		return m
	}
	c := len(m[0])
	t := make([][]int, c)
	for i := 0; i < c; i++ {
		t[i] = make([]int, r)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			t[j][i] = m[i][j]
		}
	}
	return t
}

func multiply1(A [][]int, B [][]int) [][]int {
	// b := transpose(B)
	rA, cA, _, rC := len(A), len(A[0]), len(B), len(B[0])
	r := make([][]int, rA)
	for i := 0; i < rA; i++ {
		c := make([]int, rC)
		for j := 0; j < rC; j++ {
			s := 0
			for k := 0; k < cA; k++ {
				// if A[i][k] == 0 || B[k][j] == 0 {
				// 	continue
				// }
				s += A[i][k] * B[k][j]
			}
			c[j] = s
		}
		r[i] = c
	}
	return r
}

func sparseMatrixRow(m [][]int) [][]int {
	r, c := len(m), len(m[0])
	sm := make([][]int, r)
	for i := 0; i < r; i++ {
		sr := make([]int, 0)
		for j := 0; j < c; j++ {
			if m[i][j] != 0 {
				sr = append(sr, j, m[i][j])
			}
		}
		sm[i] = sr
	}
	return sm
}

func sparseMatrixCol(m [][]int) [][]int {
	r, c := len(m), len(m[0])
	sm := make([][]int, c)
	for j := 0; j < c; j++ {
		sc := make([]int, 0)
		for i := 0; i < r; i++ {
			if m[i][j] != 0 {
				sc = append(sc, i, m[i][j])
			}
		}
		sm[j] = sc
	}
	return sm
}

func mul(a, b []int) int {
	lenA, lenB := len(a), len(b)
	aB, aS, lB, lS := a, b, lenA, lenB
	if lenB > lenA {
		lS = lenA
		lB = lenB
		aB = b
		aS = a
	}
	s := 0
	j := 0
	for i := 0; i < lS; i = i + 2 {
		index := aS[i]
		val := aS[i+1]
		for ; j < lB && j <= index; j = j + 2 {
			if index == aB[j] {
				s += val * aB[j+1]
				break
			}
		}
	}
	return s
}

func multiplySparse(A [][]int, B [][]int) [][]int {
	SA, SB := sparseMatrixRow(A), sparseMatrixCol(B)
	rA, cB := len(SA), len(SB)
	SR := make([][]int, rA)
	for i := 0; i < rA; i++ {
		c := make([]int, 0)
		for j := 0; j < cB; j++ {
			s := mul(SA[i], SB[j])
			if s != 0 {
				c = append(c, j, s)
			}
		}
		SR[i] = c
	}
	return SR
}

func multiply(A [][]int, B [][]int) [][]int {
	mA, nA, _, nB := len(A), len(A[0]), len(B), len(B[0])
	C := make([][]int, mA)
	for i := 0; i < mA; i++ {
		c := make([]int, nB)
		for k := 0; k < nA; k++ {
			if A[i][k] != 0 {
				for j := 0; j < nB; j++ {
					if B[k][j] != 0 {
						c[j] += A[i][k] * B[k][j]
					}
				}
			}
		}
		C[i] = c
	}
	return C
}

func main() {
	// A := [][]int{{1, 0, 0}, {-1, 0, 3}}
	// B := [][]int{
	// 	{7, 0, 0},
	// 	{0, 0, 0},
	// 	{0, 0, 1},
	// }
	A := [][]int{{1, -5}}
	B := [][]int{
		{12},
		{-1},
	}
	C := multiply(A, B)
	fmt.Println(C)
}
