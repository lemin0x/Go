package main

import (
	"fmt"
	"os"
)

const Size = 9

func printBoard(board [][]rune) {
	for _, row := range board {
		for i, e := range row {
			fmt.Printf("%c", e)
			if i != len(row)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func isValid(board [][]rune, row, col int, num rune) bool {
	for i := 0; i < Size; i++ {
		if board[row][i] == num || board[i][col] == num || board[row-row%3+i/3][col-col%3+i%3] == num {
			return false
		}
		
	}
	return true
}
func findEmptyCell(board [][]rune) (int, int) {
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			if board[row][col] == '.' {
				return row, col
			}
		}
	}
	return -1, -1
}
func solveSudoku(board [][]rune) bool {
	row, col := findEmptyCell(board)
	if row == -1 && col == -1 {
		return true
	}
	for num := '1'; num <= '9'; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if solveSudoku(board) {
				return true
			}
			board[row][col] = '.'
		}
	}
	return false
}
func main() {
	args := os.Args[1:]
	if len(args) != Size {
		fmt.Println("Error")
		return
	}
	board := make([][]rune, Size)
	for i := 0; i < Size; i++ {
		if len(args[i]) != Size {
			fmt.Println("Error")
			return
		}
		board[i] = []rune(args[i])
	}
	if solveSudoku(board) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}
