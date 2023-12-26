package mymath

func Power1(base float64, power int) float64 {
	sum := 1.0
	for i := 0; i < power; i++ {
		sum *= base
	}
	return sum
}

func Power2(base float64, power int, nalloc, nbytes int) float64 {
	for i := 0; i < nalloc; i++ {
		_ = make([]byte, nbytes)
	}
	return Power1(base, power)
}
