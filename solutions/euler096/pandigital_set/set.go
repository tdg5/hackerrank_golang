package pandigital_set

import (
	"fmt"
	"github.com/tdg5/hackerrank/solutions/euler096/value"
)

type PandigitalSet struct {
	Members [9]*value.Value
	size    int
}

func (set *PandigitalSet) Add(val *value.Value) {
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
