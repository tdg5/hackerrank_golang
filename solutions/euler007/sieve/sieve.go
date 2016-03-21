package sieve

import (
	"github.com/tdg5/hackerrank/solutions/euler007/window"
)

func New(windowSize uint) *sieve {
	instance := new(sieve)
	instance.window = window.New(windowSize)
	return instance
}

type sieve struct {
	window *window.Window
}

func (sieve *sieve) Ordinal(ordinal uint) uint {
	var prime uint
	if ordinal == 0 {
		return 0
	}
	if ordinal == 1 {
		return 2
	}
	for ordinal > sieve.window.Count {
		prime = sieve.window.Next()
	}
	return prime
}
