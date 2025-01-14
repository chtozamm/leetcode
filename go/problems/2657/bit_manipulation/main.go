package main

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	ans := make([]int, n)
	bitmask, common := 0, 0

	for i := 0; i < n; i++ {
		if bitmask&(1<<A[i]) != 0 {
			common++
		} else {
			bitmask |= (1 << A[i])
		}

		if bitmask&(1<<B[i]) != 0 {
			common++
		} else {
			bitmask |= (1 << B[i])
		}

		ans[i] = common
	}

	return ans
}
