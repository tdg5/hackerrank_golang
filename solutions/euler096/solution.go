package main

import (
	"fmt"
	"github.com/tdg5/hackerrank/solutions/euler096/parser"
	"github.com/tdg5/hackerrank/solutions/euler096/sudoku"
)

func main() {
	sudoku := new(sudoku.Sudoku)
	sudokuVector := parser.Parse()
	sudoku.Load(sudokuVector)
	sudoku.Resolve()
	fmt.Println(sudoku)
}
