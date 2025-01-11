package main

import "math/bits"

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

	if len(s) == k {
		return true
	}

	oddCount := 0
	for _, char := range s {
		oddCount ^= 1 << (char - 'a')
	}

	return bits.OnesCount(uint(oddCount)) <= k
}
