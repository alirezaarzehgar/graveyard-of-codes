package mymath_test

import (
	"math"
	"testing"

	mymath "github.com/alirezaarzehgar/bench"
)

func TestPower1(t *testing.T) {
	testCases := []struct {
		Base           float64
		Power          int
		ExpectedResult float64
	}{
		{Base: 1, Power: 1, ExpectedResult: 1},
		{Base: 2, Power: 2, ExpectedResult: 4},
		{Base: 3, Power: 3, ExpectedResult: 27},
		{Base: 2, Power: 10, ExpectedResult: 1024},
	}

	for _, c := range testCases {
		res := mymath.Power1(c.Base, c.Power)
		if res != c.ExpectedResult {
			t.Errorf("%f^%d == %f not %f", c.Base, c.Power, c.ExpectedResult, res)
		}

		res = mymath.Power2(c.Base, c.Power, 0, 0)
		if res != c.ExpectedResult {
			t.Errorf("%f^%d == %f not %f", c.Base, c.Power, c.ExpectedResult, res)
		}
	}
}

var result float64

func BenchmarkPower1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := mymath.Power1(2, 4)
		result = res
	}
}

func BenchmarkStdPower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := math.Pow(2, 4)
		result = res
	}
}

func pAllocBench(b *testing.B, nalloc, nbytes int) {
	for i := 0; i < b.N; i++ {
		res := mymath.Power2(2, 4, nalloc, nbytes)
		result = res
	}
}
func BenchmarkPowerAlloc_1_1(b *testing.B) {
	pAllocBench(b, 1, 1)
}

func BenchmarkPowerAlloc_2_1(b *testing.B) {
	pAllocBench(b, 2, 1)
}

func BenchmarkPowerAlloc_3_1(b *testing.B) {
	pAllocBench(b, 3, 1)
}

func BenchmarkPowerAlloc_1_1024(b *testing.B) {
	pAllocBench(b, 1, 1024)
}

func BenchmarkPowerAlloc_1_2048(b *testing.B) {
	pAllocBench(b, 1, 2048)
}

func BenchmarkPowerAlloc_1_4056(b *testing.B) {
	pAllocBench(b, 1, 4056)
}

func BenchmarkPowerAlloc_2_2048(b *testing.B) {
	pAllocBench(b, 2, 2048)
}

func BenchmarkPowerAlloc_3_4056(b *testing.B) {
	pAllocBench(b, 3, 4056)
}
