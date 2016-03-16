package main

import (
	"bufio"
	"fmt"
	"github.com/tdg5/hackerrank/solutions/euler007/ordinator"
	"github.com/tdg5/hackerrank/solutions/euler007/sieve"
	"os"
	"strconv"
)

func main() {
	var ordinator ordinator.Ordinator
	ordinator = sieve.New(0)
	scanner := bufio.NewScanner(os.Stdin)
	// Scan once to skip the test case count
	scanner.Scan()

	for scanner.Scan() {
		ordinal, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		prime := ordinator.Ordinal(uint(ordinal))
		fmt.Println(prime)
	}
}
