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

func (value *Value) String() string {
	return fmt.Sprintf("%d", value.Value)
}

type ValueSet interface {
	Add(*Value)
	Resolve() bool
	String() string
	UnknownCount() int
}
