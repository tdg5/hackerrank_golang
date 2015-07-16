package pandigital_set

import(
  "fmt"
)

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
