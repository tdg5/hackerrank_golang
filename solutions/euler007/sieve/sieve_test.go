package sieve

import (
	"testing"
)

func TestNew(t *testing.T) {
	subject := New(0)
	if subject.Ordinal(1) != 2 {
		t.Error("wah")
	}
}
