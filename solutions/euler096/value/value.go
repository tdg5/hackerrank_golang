package value

import (
	"fmt"
)

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
