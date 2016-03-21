package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Ordinator interface {
	Ordinal(ordinal uint) uint
}

const DefaultWindowSize = 10000

type sieve struct {
	cursorIndex  uint
	primes       []uint
	primeCount   uint
	window       []bool
	windowOffset uint
	windowSize   uint
	windowSpan   uint
}

func New(windowSize uint) *sieve {
	instance := new(sieve)
	if windowSize != 0 {
		instance.windowSize = windowSize
	} else {
		instance.windowSize = DefaultWindowSize
	}
	instance.initialize()
	return instance
}

func (sieve *sieve) Ordinal(ordinal uint) uint {
	if ordinal == 0 {
		return 0
	}
	for ordinal > sieve.primeCount {
		sieve.findNext()
	}
	return sieve.primes[ordinal-1]
}

func (sieve *sieve) advanceCursor() {
	if sieve.cursorIndex >= sieve.windowSize-1 {
		sieve.shiftWindow(sieve.windowOffset + sieve.windowSpan)
		return
	}

	sieve.cursorIndex++
}

func (sieve *sieve) applyKnownPrimes() {
	for i := uint(1); i < sieve.primeCount; i++ {
		sieve.applyPrime(sieve.primes[i])
	}
}

func (sieve *sieve) applyPrime(prime uint) {
	cursorOffset := sieve.cursorValue()
	firstRemainder := cursorOffset % (prime * 2)
	applicationValue := cursorOffset - firstRemainder + prime
	if firstRemainder > prime {
		applicationValue += 2 * prime
	}
	applicationIndex := (applicationValue - sieve.windowOffset) / 2
	for applicationIndex < sieve.windowSize {
		sieve.window[applicationIndex] = false
		applicationIndex += prime
	}
}

func (sieve *sieve) cursorValue() uint {
	return sieve.windowOffset + (2 * sieve.cursorIndex)
}

func (sieve *sieve) findNext() {
	for !sieve.window[sieve.cursorIndex] {
		sieve.advanceCursor()
	}
	prime := sieve.cursorValue()
	sieve.primes = append(sieve.primes, prime)
	sieve.primeCount++
	sieve.applyPrime(prime)
	sieve.advanceCursor()
}

func (sieve *sieve) initialize() {
	sieve.primes = []uint{2}
	sieve.window = make([]bool, sieve.windowSize)
	sieve.windowSpan = sieve.windowSize * 2
	sieve.primeCount = 1
	sieve.shiftWindow(3)
}

func (sieve *sieve) shiftWindow(offset uint) {
	for i := uint(0); i < sieve.windowSize; i++ {
		sieve.window[i] = true
	}
	sieve.windowOffset = offset
	sieve.cursorIndex = uint(0)
	sieve.applyKnownPrimes()
}

func main() {
	var ordinator Ordinator
	ordinator = New(0)
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
