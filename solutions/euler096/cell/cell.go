package cell

import(
  "github.com/tdg5/hackerrank/solutions/euler096/pandigital_set"
)

type Cell struct {
  Column *pandigital_set.PandigitalSet
  Row *pandigital_set.PandigitalSet
  Square *pandigital_set.PandigitalSet
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
