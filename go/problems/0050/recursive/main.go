package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	if n == 0 {
		return 1.0
	}

	half := myPow(x, n/2)

	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}
