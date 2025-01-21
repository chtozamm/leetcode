package main

import "math"

func gridGame(grid [][]int) int64 {
	firstRowSum := 0
	for _, num := range grid[0] {
		firstRowSum += num
	}

	secondRowSum := 0
	minimumSum := math.MaxInt

	for i, num := range grid[0] {
		firstRowSum -= num
		minimumSum = min(minimumSum, max(firstRowSum, secondRowSum))
		secondRowSum += grid[1][i]
	}

	return int64(minimumSum)
}
