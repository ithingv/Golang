package main

import (
	"fmt"
	"strings"
)

func main() {
	a := make([]int, 5)    // len(5)
	b := make([]int, 0, 5) // len(b)=0, cap(b) = 5

	b = b[:cap(b)] // len(b)=5, cap(b)=5
	fmt.Println(a)
	fmt.Println(b)

	board := make_board()
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}

func make_board() [][]string {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	return board
}
