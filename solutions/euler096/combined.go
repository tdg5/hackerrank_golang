package euler096

import(
  "fmt"
  "bufio"
  "os"
  "strconv"
)

type Sudoku struct {
  values [81] int
  cells [81] Cell
  rows [9] PandigitalSet
  columns [9] PandigitalSet
  squares [9] PandigitalSet
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
     sudoku.cells[offset].Value = &sudoku.values[offset]
  }
  sudoku.link()
}

func (sudoku *Sudoku) Resolve() {
  for sudoku.UnknownCount() != 0 {
    for _, cell := range sudoku.cells {
      cell.Resolve()
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
      valueIndex := yOffset * 9 + xOffset
      column := &sudoku.columns[xOffset]
      column.Add(&sudoku.values[valueIndex])
      sudoku.cells[valueIndex].Column = column
    }
  }
}

func (sudoku *Sudoku) linkRows() {
  for yOffset := 0; yOffset < 9; yOffset++ {
    for xOffset := 0; xOffset < 9; xOffset++ {
      valueIndex := xOffset * 9 + yOffset
      row := &sudoku.rows[xOffset]
      row.Add(&sudoku.values[valueIndex])
      sudoku.cells[valueIndex].Row = row
    }
  }
}

func (sudoku *Sudoku) linkSquares() {
  for index := 0; index < 81; index++ {
    xOffset, yOffset := index / 9, index % 9
    squaredIndex := ((xOffset % 3) * 3) + ((xOffset / 3) * 27) + ((yOffset / 3) * 9) + (yOffset % 3)
    squareIndex := squaredIndex / 9
    square := &sudoku.squares[squareIndex]
    square.Add(&sudoku.values[index])
    sudoku.cells[index].Square = square
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
type PandigitalSet struct {
  Members [9] *int
  size int
}

func (set *PandigitalSet) Add(value *int) {
  set.Members[set.size] = value
  set.size++
}

func (set PandigitalSet) String() string {
  str := ""
  for _, value := range set.Members {
    str += fmt.Sprintf("%d", *value)
  }
  return str
}

func (set *PandigitalSet) UnknownCount() int {
  count := 0
  for _, value := range set.Members {
    if *value == 0 { count++ }
  }
  return count
}

func (set *PandigitalSet) KnownDigitMask() int {
  var known int
  for _, value := range set.Members {
    if *value == 0 { continue }
    known = known | (1 << uint(*value))
  }
  return known
}

func (set *PandigitalSet) Resolve() {
  if set.UnknownCount() != 1 { return }

  missingDigit := 45
  var zeroIndex int
  for index, value := range set.Members {
    if *value == 0 {
      zeroIndex = index
    } else {
      missingDigit -= *value
    }
  }
  *set.Members[zeroIndex] = missingDigit
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

type Cell struct {
  Column *PandigitalSet
  Row *PandigitalSet
  Square *PandigitalSet
  Value *int
}

func (cell *Cell) Resolve() {
  if *cell.Value != 0 { return }
  unavailable := cell.Row.KnownDigitMask() | cell.Column.KnownDigitMask() | cell.Square.KnownDigitMask()
  availableCount := 0
  solo := 0
  for val := 1; val < 10; val++ {
    bin := 1 << uint(val)
    if unavailable & bin != bin {
      availableCount++
      solo = val
    }
  }
  if availableCount != 1 { return }
  *cell.Value = solo
  cell.Row.Resolve()
  cell.Column.Resolve()
  cell.Square.Resolve()
}
