package parser

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

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
