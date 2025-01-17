package main

func doesValidArrayExist(derived []int) bool {
	sum := 0

	for _, val := range derived {
		sum += val
	}

	return sum%2 == 0
}
