package sieve

import (
	"testing"
)

func TestNewWithDefaultWindowSize(t *testing.T) {
	subject := *New(0)
	if subject.windowSize != DefaultWindowSize {
		t.Errorf("sieve.New(0).windowSize != %i", DefaultWindowSize)
	}
}

func TestNewWithExplicitWindowSize(t *testing.T) {
	windowSize := uint(100)
	subject := *New(windowSize)
	if subject.windowSize != windowSize {
		t.Errorf("sieve.New(%d).windowSize != %d", windowSize, windowSize)
	}
}

func TestNewInitialization(t *testing.T) {
	subject := *New(0)
	if subject.primes[0] != 2 {
		t.Error("sieve.New().primeSlice[0] != 2")
	}
	if subject.primeCount != 1 {
		t.Error("sieve.New().primeCount != 1")
	}
	if subject.windowOffset != 3 {
		t.Error("sieve.New().windowOffset != 3")
	}
	if !subject.window.Test(0) {
		t.Error("sieve.New().window.Test(0) == false")
	}
	if subject.window.Len() != subject.windowSize {
		t.Error("sieve.New(windowSize).window.Len() != windowSize")
	}
	if subject.cursorIndex != 0 {
		t.Error("sieve.New().cursorIndex != 0")
	}
	if subject.windowSpan != subject.windowSize*2 {
		t.Error("sieve.New(windowSpan).windowSpan != windowSpan * 2")
	}
}

func TestOrdinalOf0(t *testing.T) {
	subject := *New(0)
	ord := subject.Ordinal(0)
	if ord != 0 {
		t.Errorf("sieve.Ordinal(0) == %d, expected 0", ord)
	}
}

func TestOrdinalOf1(t *testing.T) {
	subject := *New(0)
	ord := subject.Ordinal(1)
	if ord != 2 {
		t.Errorf("sieve.Ordinal(1) == %d, expected 2", ord)
	}
}

func TestOrdinalOf2(t *testing.T) {
	subject := *New(0)
	ord := subject.Ordinal(2)
	if ord != 3 {
		t.Errorf("sieve.Ordinal(2) == %d, expected 3", ord)
	}
}

func TestOrdinalOf5(t *testing.T) {
	subject := *New(2)
	ord := subject.Ordinal(5)
	expected := uint(11)
	if ord != expected {
		t.Errorf("sieve.Ordinal(5) == %d, expected %d", ord, expected)
	}
}

func TestOrdinalOf10001(t *testing.T) {
	subject := *New(0)
	ord := subject.Ordinal(10001)
	expected := uint(104743)
	if ord != expected {
		t.Errorf("sieve.Ordinal(10001) == %d, expected %d", ord, expected)
	}
}

func TestCursorValue(t *testing.T) {
	subject := *New(0)
	cursorVal := subject.cursorValue()
	if cursorVal != 3 {
		t.Errorf("sieve.New().cursorValue() == %d, expected 3", cursorVal)
	}
}
