package sudoku

import (
	"fmt"
	"github.com/tdg5/hackerrank/solutions/euler096/pandigital_set"
	"github.com/tdg5/hackerrank/solutions/euler096/value"
)

type Sudoku struct {
	values  [81]*value.Value
	rows    [9]pandigital_set.PandigitalSet
	columns [9]pandigital_set.PandigitalSet
	sets    [27]*pandigital_set.PandigitalSet
	squares [9]pandigital_set.PandigitalSet
}

func (sudoku *Sudoku) Inspect() {
	sudoku.PrintValues()
	sudoku.PrintRows()
	sudoku.PrintColumns()
	sudoku.PrintSquares()
}

func (sudoku *Sudoku) Load(values [81]int) {
	for offset := 0; offset < 81; offset++ {
		sudoku.values[offset] = &value.Value{Value: values[offset]}
	}
	sudoku.link()
}

func (sudoku *Sudoku) Resolve() {
	for sudoku.UnknownCount() != 0 {
		for index := 0; index < len(sudoku.sets); index++ {
			sudoku.sets[index].Resolve()
		}
	}
}

func (sudoku *Sudoku) UnknownCount() int {
	count := 0
	for index := 0; index < len(sudoku.values); index++ {
		if sudoku.values[index].Value == 0 {
			count++
		}
	}
	return count
}

func (sudoku *Sudoku) link() {
	sudoku.linkRows()
	sudoku.linkColumns()
	sudoku.linkSquares()
	sudoku.trackSets()
}

func (sudoku *Sudoku) linkColumns() {
	for yOffset := 0; yOffset < 9; yOffset++ {
		for xOffset := 0; xOffset < 9; xOffset++ {
			column := &sudoku.columns[xOffset]
			val := sudoku.values[yOffset*9+xOffset]
			column.Add(val)
			var valueSet value.ValueSet = column
			val.Column = &valueSet
		}
	}
}

func (sudoku *Sudoku) linkRows() {
	for yOffset := 0; yOffset < 9; yOffset++ {
		for xOffset := 0; xOffset < 9; xOffset++ {
			row := &sudoku.rows[xOffset]
			val := sudoku.values[xOffset*9+yOffset]
			row.Add(val)
			var valueSet value.ValueSet = row
			val.Row = &valueSet
		}
	}
}

func (sudoku *Sudoku) linkSquares() {
	for index := 0; index < 81; index++ {
		xOffset, yOffset := index/9, index%9
		squaredIndex := ((xOffset % 3) * 3) + ((xOffset / 3) * 27) + ((yOffset / 3) * 9) + (yOffset % 3)
		squareIndex := squaredIndex / 9
		square := &sudoku.squares[squareIndex]
		val := sudoku.values[index]
		square.Add(val)
		var valueSet value.ValueSet = square
		val.Square = &valueSet
	}
}

func (sudoku *Sudoku) PrintColumns() {
	str := "Columns:"
	for xOffset := 0; xOffset < 9; xOffset++ {
		str += fmt.Sprintf("\n\t%v", &sudoku.columns[xOffset])
	}
	fmt.Println(str)
}

func (sudoku *Sudoku) PrintRows() {
	fmt.Println("Rows:", sudoku)
}

func (sudoku *Sudoku) PrintSquares() {
	str := "Squares"
	for _, value := range sudoku.squares {
		str += fmt.Sprintf("\n\t%v", &value)
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
		if yOffset != 0 {
			str += "\n"
		}
		str += fmt.Sprintf("%v", &sudoku.rows[yOffset])
	}
	return str
}

func (sudoku *Sudoku) trackSets() {
	index := 0
	for rowIndex := 0; rowIndex < len(sudoku.rows); rowIndex++ {
		sudoku.sets[index] = &sudoku.rows[rowIndex]
		index++
	}
	for columnIndex := 0; columnIndex < len(sudoku.columns); columnIndex++ {
		sudoku.sets[index] = &sudoku.columns[columnIndex]
		index++
	}
	for squareIndex := 0; squareIndex < len(sudoku.squares); squareIndex++ {
		sudoku.sets[index] = &sudoku.squares[squareIndex]
		index++
	}
}
