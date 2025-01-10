package main

func myPow(x float64, n int) float64 {
	res := 1.0

	if x == 1.0 {
		return 1.0
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	for n > 0 {
		if (n & 1) != 0 {
			res *= x
		}
		n = n >> 1
		x *= x
	}

	return res
}
