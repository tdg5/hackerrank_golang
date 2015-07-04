package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Scan once to skip the test case count
	scanner.Scan()

	for scanner.Scan() {
		limit, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		// Subtract 1 since we interested in numbers under the limit
		limit -= 1

		sum := binomial(limit/3) * 3
		sum += binomial(limit/5) * 5
		sum -= binomial(limit/15) * 15

		fmt.Println(sum)
	}
}

func binomial(limit int) int {
	return limit * (limit + 1) / 2
}
