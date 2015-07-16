package main

import (
  "fmt"
	"github.com/tdg5/hackerrank/solutions/euler096/parser"
	"github.com/tdg5/hackerrank/solutions/euler096/sudoku"
)

func main() {
  sudoku := new(sudoku.Sudoku)
  sudoku.Load(parser.Parse())
  sudoku.Resolve()
  fmt.Println(sudoku)
}
