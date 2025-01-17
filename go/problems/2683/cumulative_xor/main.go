package main

func doesValidArrayExist(derived []int) bool {
	xor := 0

	for _, elem := range derived {
		xor ^= elem
	}

	return xor == 0
}
