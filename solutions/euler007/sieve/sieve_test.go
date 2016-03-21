package sieve

import (
	"testing"
)

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

func BenchmarkOrdinalOf10001(b *testing.B) {
	for i := 0; i < b.N; i++ {
		subject := *New(0)
		subject.Ordinal(10001)
	}
}
