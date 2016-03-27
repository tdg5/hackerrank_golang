package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sudoku := new(Sudoku)
	sudokuVector := Parse()
	sudoku.Load(sudokuVector)
	sudoku.Resolve()
	fmt.Println(sudoku)
}

func Parse() (values [81]int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)
	var output [81]int
	count := 0
	for scanner.Scan() {
		if scanner.Text() == "\n" {
			continue
		}

		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		output[count] = value
		count++
	}
	return output
}

type Value struct {
	Column *ValueSet
	Row    *ValueSet
	Square *ValueSet
	Value  int
}

func (value *Value) availableDigits() []int {
	digitCounts := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
		9: 0,
	}
	rowDigits := (*value.Row).MissingDigits()
	for index := 0; index < len(rowDigits); index++ {
		digitCounts[rowDigits[index]]++
	}
	columnDigits := (*value.Column).MissingDigits()
	for index := 0; index < len(columnDigits); index++ {
		digitCounts[columnDigits[index]]++
	}
	squareDigits := (*value.Square).MissingDigits()
	for index := 0; index < len(squareDigits); index++ {
		digitCounts[squareDigits[index]]++
	}
	commonMissingDigits := make([]int, 0)
	for digit, count := range digitCounts {
		if count == 3 {
			commonMissingDigits = append(commonMissingDigits, digit)
		}
	}
	return commonMissingDigits
}

func (value *Value) String() string {
	return fmt.Sprintf("%d", value.Value)
}

func (value *Value) Resolve() bool {
	if value.Value != 0 {
		return false
	}
	possibleValues := value.availableDigits()
	if len(possibleValues) != 1 {
		return false
	}
	value.Value = possibleValues[0]
	return true
}

type ValueSet interface {
	Add(*Value)
	MissingDigits() []int
	Resolve() bool
	String() string
	UnknownCount() int
}

type PandigitalSet struct {
	Members [9]*Value
	size    int
}

func (set *PandigitalSet) Add(val *Value) {
	set.Members[set.size] = val
	set.size++
}

func (set *PandigitalSet) Resolve() bool {
	if set.UnknownCount() != 1 {
		return false
	}

	missingDigit := 45
	var zeroIndex int
	for index := 0; index < set.size; index++ {
		val := set.Members[index]
		if val.Value == 0 {
			zeroIndex = index
		} else {
			missingDigit -= val.Value
		}
	}
	set.Members[zeroIndex].Value = missingDigit
	return true
}

func (set *PandigitalSet) MissingDigits() []int {
	digits := [9]bool{true, true, true, true, true, true, true, true, true}
	for index := 0; index < set.size; index++ {
		val := set.Members[index].Value
		if val == 0 {
			continue
		}
		digits[val-1] = false
	}
	missingDigits := make([]int, 0)
	for index := 0; index < len(digits); index++ {
		if digits[index] == true {
			missingDigits = append(missingDigits, index+1)
		}
	}
	return missingDigits
}

func (set *PandigitalSet) String() string {
	str := ""
	for index := 0; index < set.size; index++ {
		val := set.Members[index]
		str += fmt.Sprintf("%d", val.Value)
	}
	return str
}

func (set *PandigitalSet) UnknownCount() int {
	count := 0
	for index := 0; index < set.size; index++ {
		val := set.Members[index]
		if val.Value == 0 {
			count++
		}
	}
	return count
}

type Sudoku struct {
	values  [81]*Value
	rows    [9]PandigitalSet
	columns [9]PandigitalSet
	sets    [27]*PandigitalSet
	squares [9]PandigitalSet
}

func (sudoku *Sudoku) Inspect() {
	sudoku.PrintValues()
	sudoku.PrintRows()
	sudoku.PrintColumns()
	sudoku.PrintSquares()
}

func (sudoku *Sudoku) Load(values [81]int) {
	for offset := 0; offset < 81; offset++ {
		sudoku.values[offset] = &Value{Value: values[offset]}
	}
	sudoku.link()
}

func (sudoku *Sudoku) Resolve() {
	for sudoku.UnknownCount() != 0 {
		for index := 0; index < len(sudoku.values); index++ {
			sudoku.values[index].Resolve()
		}
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
			var valueSet ValueSet = column
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
			var valueSet ValueSet = row
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
		var valueSet ValueSet = square
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
