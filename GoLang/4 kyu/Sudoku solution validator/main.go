package main

import (
	"fmt"
	"sort"
)

func main() {
	var testFalse = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 0, 3, 4, 8},
		{1, 0, 0, 3, 4, 2, 5, 6, 0},
		{8, 5, 9, 7, 6, 1, 0, 2, 0},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 0, 1, 5, 3, 7, 2, 1, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 0, 0, 4, 8, 1, 1, 7, 9},
    }
	var testTrue = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
    }
	
	fmt.Println(ValidateSolution(testTrue))
	fmt.Println(ValidateSolution(testFalse))
}

func ValidateSolution(board [][]int) bool {
	if validLines(board) && validColumns(board) && validSquares(board) {
		return true
	}
	return false
}

func validArray(row []int) bool {
	cRow := make([]int, len(row))
	copy(cRow, row) //copy row to not to mess up an actual board ("sort" sorts rows directly on the board even without a pointer)
	sort.Ints(cRow)

	if equal(cRow, []int{1,2,3,4,5,6,7,8,9}) {
		return true
	}
	return false
}

func equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }

    for index, value := range a {
        if value != b[index] {
            return false
        }
    }
    return true
}

//checks every line (horizontal)
func validLines(board [][]int) bool {
	for _, row := range board {
		if !validArray(row) {
			return false
		}
	}
	return true
}

//checks every column
func validColumns(board [][]int) bool {
	for i := 0; i < 9; i++ {
		column := []int{}
		for index := 0; index < len(board); index++ {
			column = append(column, board[index][i])
		}

		if !validArray(column) {
			return false
		}
	}
	return true
}

//checks every square in the board
func validSquares(board [][]int) bool {
	for i := 0; i < 7; i += 3 {
		for j := 0; j < 7; j += 3 {
			if !validSquare(board, i, j) {
				return false
			}
		}
	}
	return true
}

func validSquare(board [][]int, a, b int) bool {
	square := []int{}
	for x := a; x < a + 3; x++ {
		for y := b; y < b + 3; y++ {
			square = append(square, board[x][y])
		}
	}
	return validArray(square)
}