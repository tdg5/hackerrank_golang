package window

import (
	"github.com/willf/bitset"
)

const DefaultSize = 10000

type Window struct {
	Count       uint
	cursorIndex uint
	offset      uint
	primes      []uint
	segment     *bitset.BitSet
	size        uint
	span        uint
}

func New(size uint) *Window {
	instance := new(Window)
	if size != 0 {
		instance.size = size
	} else {
		instance.size = DefaultSize
	}
	instance.initialize()
	return instance
}

func (window *Window) Next() uint {
	for !window.segment.Test(window.cursorIndex) {
		window.advanceCursor()
	}
	prime := window.cursorValue()
	window.primes = append(window.primes, prime)
	window.Count++
	window.applyPrime(prime)
	window.advanceCursor()
	return prime
}

func (window *Window) advanceCursor() {
	if window.cursorIndex >= window.size-1 {
		window.shift(window.offset + window.span)
		return
	}
	window.cursorIndex++
}

func (window *Window) applyKnownPrimes() {
	for i := uint(1); i < window.Count; i++ {
		window.applyPrime(window.primes[i])
	}
}

func (window *Window) applyPrime(prime uint) {
	cursorOffset := window.cursorValue()
	firstRemainder := cursorOffset % (prime * 2)
	applicationValue := cursorOffset + prime - firstRemainder
	if firstRemainder > prime {
		applicationValue += 2 * prime
	}
	applicationIndex := (applicationValue - window.offset) / 2
	for applicationIndex < window.size {
		window.segment.Clear(applicationIndex)
		applicationIndex += prime
	}
}

func (window *Window) cursorValue() uint {
	return window.offset + (2 * window.cursorIndex)
}

func (window *Window) initialize() {
	window.primes = []uint{2}
	window.segment = bitset.New(window.size)
	window.span = window.size * 2
	window.shift(3)
	window.Count = 1
}

func (window *Window) shift(offset uint) {
	for i := uint(0); i < window.size; i++ {
		window.segment.Set(i)
	}
	window.offset = offset
	window.cursorIndex = uint(0)
	window.applyKnownPrimes()
}
