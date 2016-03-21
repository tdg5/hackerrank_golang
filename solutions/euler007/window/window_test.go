package window

import (
	"testing"
)

func TestNewWithDefaultWindowSize(t *testing.T) {
	subject := *New(0)
	if subject.size != DefaultSize {
		t.Errorf("window.New(0).size != %i", DefaultSize)
	}
}

func TestNewWithExplicitWindowSize(t *testing.T) {
	size := uint(100)
	subject := *New(size)
	if subject.size != size {
		t.Errorf("window.New(%d).size != %d", size, size)
	}
}

func TestNewInitialization(t *testing.T) {
	subject := *New(0)
	if subject.primes[0] != 2 {
		t.Error("window.New().primes[0] != 2")
	}
	if subject.Count != 1 {
		t.Error("window.New().Count != 1")
	}
	if subject.offset != 3 {
		t.Error("window.New().offset != 3")
	}
	if !subject.segment.Test(0) {
		t.Error("window.New().segment.Test(0) == false")
	}
	if subject.segment.Len() != subject.size {
		t.Error("window.New(size).segment.Len() != size")
	}
	if subject.cursorIndex != 0 {
		t.Error("window.New().cursorIndex != 0")
	}
	if subject.span != subject.size*2 {
		t.Error("window.New(size).span != size * 2")
	}
}

func TestCursorValue(t *testing.T) {
	subject := *New(0)
	cursorVal := subject.cursorValue()
	if cursorVal != 3 {
		t.Errorf("window.New().cursorValue() == %d, expected 3", cursorVal)
	}
}
