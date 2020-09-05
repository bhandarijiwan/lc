// https://leetcode.com/problems/design-tic-tac-toe/

package main

import (
	"fmt"
)

type TicTacToe struct {
	board  [][]int
	winner int
}

func (t *TicTacToe) String() string {
	s := ""
	for _, row := range t.board {
		for i, cell := range row {
			x := " "
			if cell == 1 {
				x = "x"
			} else if cell == 2 {
				x = "o"
			}
			y := ""
			if i == 0 {
				y = "|"
			}
			s = s + y + x + "|"
		}
		s = s + "\n"
	}
	s = s + ""
	return s
}

/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}
	return TicTacToe{board: board}
}

func (this *TicTacToe) CheckWinner(row, col, player int) {
	// along the row
	i, l := 0, len(this.board)
	for ; i < l; i++ {
		if this.board[row][i] != player {
			break
		}
	}
	if i >= l {
		this.winner = player
		return
	}
	// along the column
	i = 0
	for ; i < l; i++ {
		if this.board[i][col] != player {
			break
		}
	}
	if i >= l {
		this.winner = player
		return
	}
	// along the \ diagonal
	i = 0
	if row == col {
		for ; i < l; i++ {
			if this.board[i][i] != player {
				break
			}
		}
		if i >= l {
			this.winner = player
			return
		}
	}
	// along the / diagonal
	i = 0
	k := l - 1
	if k-col == row && k-row == col {
		for ; i < l; i++ {
			if this.board[i][k-i] != player {
				break
			}
		}
		if i >= l {
			this.winner = player
		}
	}
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
	if this.winner >= 1 {
		return this.winner
	}
	this.board[row][col] = player
	this.CheckWinner(row, col, player)
	return this.winner
}

func main() {
	game := Constructor(3)
	fmt.Println(game.Move(0, 0, 1))
	fmt.Println(game.Move(0, 1, 1))
	fmt.Println(game.Move(0, 2, 1))


	// fmt.Println(game.Move(2, 0, 2))
	// fmt.Println(game.Move(1, 0, 2))
	// fmt.Println(game.Move(2, 1, 1))
}
