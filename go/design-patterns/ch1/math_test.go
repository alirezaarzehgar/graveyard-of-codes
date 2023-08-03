package ch1

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	if multiply(5, 7) != 35 {
		t.Errorf("Failed")
	}
}

func TestSum(t *testing.T) {
	if sum(2, 3) != 5 {
		t.Error("Failed")
	}
}
func TestSub(t *testing.T) {
	if sub(5, 3) != 2 {
		t.Error("Failed")
	}
}

func TestDiv(t *testing.T) {
	if div(10, 2) != 5.0 {
		t.Error("Failed")
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < 100000000; i++ {
		multiply(3, i)
	}
}
