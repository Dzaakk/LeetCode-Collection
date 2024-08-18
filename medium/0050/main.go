func myPow(x float64, n int) float64 {
	pow := float64(1)
	a := x
	min := false

	if n < 0 {
		min = true
		n = -n
	}

	for n > 0 {
		if (n & 1) != 0 {
			pow = pow * a
		}

		a = a * a
		n >>= 1
	}
	if min {
		pow = 1 / pow
	}

	return pow
}