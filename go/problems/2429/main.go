package main

import (
	"math/bits"
)

func minimizeXor(num1 int, num2 int) int {
	res := num1

	targetSetBitsCount := bits.OnesCount(uint(num2))
	setBitsCount := bits.OnesCount(uint(res))

	currentBit := 0
	for setBitsCount < targetSetBitsCount {
		if res&(1<<currentBit) == 0 {
			res |= 1 << currentBit
			setBitsCount++
		}
		currentBit++
	}

	for setBitsCount > targetSetBitsCount {
		if res&(1<<currentBit) != 0 {
			res &= ^(1 << currentBit)
			setBitsCount--
		}
		currentBit++
	}

	return res
}
