package sudoku

import (
  "fmt"
  "github.com/tdg5/hackerrank/solutions/euler096/pandigital_set"
)

type Sudoku struct {
  values [81] int
  rows [9] pandigital_set.PandigitalSet
  columns [9] pandigital_set.PandigitalSet
  squares [9] pandigital_set.PandigitalSet
}

func (sudoku *Sudoku) Inspect() {
  sudoku.PrintValues()
  sudoku.PrintRows()
  sudoku.PrintColumns()
  sudoku.PrintSquares()
}

func (sudoku *Sudoku) Load(values [81]int) {
  for offset := 0; offset < 81; offset++ {
    sudoku.values[offset] = values[offset]
  }
  sudoku.link()
}

func (sudoku *Sudoku) Resolve() {
  for sudoku.UnknownCount() != 0 {
    for _, row := range sudoku.rows {
      row.Resolve()
    }
    for _, column := range sudoku.columns {
      column.Resolve()
    }
    for _, square := range sudoku.squares {
      square.Resolve()
    }
  }
}

func (sudoku *Sudoku) UnknownCount() int {
  count := 0
  for _, value := range sudoku.values {
    if value == 0 { count++ }
  }
  return count
}

func (sudoku *Sudoku) link() {
  sudoku.linkRows()
  sudoku.linkColumns()
  sudoku.linkSquares()
}

func (sudoku *Sudoku) linkColumns() {
  for yOffset := 0; yOffset < 9; yOffset++ {
    for xOffset := 0; xOffset < 9; xOffset++ {
      sudoku.columns[xOffset].Add(&sudoku.values[yOffset * 9 + xOffset])
    }
  }
}

func (sudoku *Sudoku) linkRows() {
  for yOffset := 0; yOffset < 9; yOffset++ {
    for xOffset := 0; xOffset < 9; xOffset++ {
      sudoku.rows[xOffset].Add(&sudoku.values[xOffset * 9 + yOffset])
    }
  }
}

func (sudoku *Sudoku) linkSquares() {
  for index := 0; index < 81; index++ {
    xOffset, yOffset := index / 9, index % 9
    squaredIndex := ((xOffset % 3) * 3) + ((xOffset / 3) * 27) + ((yOffset / 3) * 9) + (yOffset % 3)
    squareIndex := squaredIndex / 9
    sudoku.squares[squareIndex].Add(&sudoku.values[index])
  }
}

func (sudoku *Sudoku) PrintColumns() {
  str := "Columns:"
  for xOffset := 0; xOffset < 9; xOffset++ {
    str += fmt.Sprintf("\n\t%v", sudoku.columns[xOffset])
  }
  fmt.Println(str)
}

func (sudoku *Sudoku) PrintRows() {
  fmt.Println("Rows:", sudoku)
}

func (sudoku *Sudoku) PrintSquares() {
  str := "Squares"
  for _, value := range sudoku.squares {
    str += fmt.Sprintf("\n\t%v", value)
  }
  fmt.Println(str)
}

func (sudoku *Sudoku) PrintValues() {
  str := "Values:\n\t"
  for _, value := range sudoku.values {
    str += fmt.Sprintf("%d ", value)
  }
  fmt.Println(str)
}

func (sudoku Sudoku) String() string {
  str := ""
  for yOffset := 0; yOffset < 9; yOffset++ {
    if yOffset != 0 { str += "\n" }
    str += fmt.Sprintf("%v", sudoku.rows[yOffset])
  }
  return str
}
