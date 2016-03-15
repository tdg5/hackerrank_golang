package sieve

import (
	"github.com/willf/bitset"
)

const DefaultWindowSize = 10000
const DefaultPrimeSliceSize = 1000

func New(windowSize uint) *sieve {
	sieve := new(sieve)
	if windowSize != 0 {
		sieve.windowSize = windowSize
	} else {
		sieve.windowSize = DefaultWindowSize
	}
	sieve.initialize()
	return sieve
}

type sieve struct {
	cursorIndex  uint
	primes       *[]uint
	primeCount   uint
	window       *bitset.BitSet
	windowOffset uint
	windowSize   uint
}

func (sieve *sieve) Ordinal(ordinal uint) uint {
	if ordinal > uint(len(*sieve.primes)+1) {
		sieve.scanUntil(ordinal)
	}
	return (*sieve.primes)[ordinal]
}

func (sieve *sieve) initialize() {
	primeSlice := make([]uint, DefaultPrimeSliceSize)
	sieve.primes = &primeSlice
	sieve.window = bitset.New(sieve.windowSize)
	sieve.shiftWindow(2)
	primeSlice[0] = 2
	sieve.primeCount = uint(1)
}

func (sieve *sieve) scanUntil(ordinal uint) {
	for ordinal > sieve.primeCount {
		sieve.checkNext()
	}
}

func (sieve *sieve) checkNext() {
	for sieve.windowOffset < uint(sieve.window.Len()) {
		if !sieve.window.Test(sieve.windowOffset) {
			sieve.windowOffset++
			continue
		}
		(*sieve.primes)[sieve.primeCount] = uint(sieve.windowOffset)
		sieve.primeCount++
	}
}

func (sieve *sieve) shiftWindow(offset uint) {
	for i := uint(0); i < sieve.windowSize; i++ {
		sieve.window.Set(i)
	}
	sieve.windowOffset = offset
	sieve.cursorIndex = uint(0)
	sieve.applyKnownPrimes()
}

func (sieve *sieve) advanceCursor() {
	if sieve.cursorIndex >= sieve.windowSize {
		sieve.shiftWindow(sieve.windowOffset + sieve.windowSize)
		return
	}

	sieve.cursorIndex++
	sieve.windowOffset++
}

func (sieve *sieve) applyKnownPrimes() {
	for i := uint(0); i < sieve.primeCount; i++ {
		prime := (*sieve.primes)[i]
		firstRemainder := sieve.windowOffset % prime
		firstInWindow := sieve.windowOffset - firstRemainder
		if firstRemainder > 0 {
			firstInWindow = firstInWindow + prime
		}
		firstIndex := firstInWindow - sieve.windowOffset
		for j := firstIndex; j < sieve.windowSize; j = j + prime {
			sieve.window.Clear(j)
		}
	}
}
