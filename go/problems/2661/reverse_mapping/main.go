package main

import "math"

func firstCompleteIndex(arr []int, mat [][]int) int {
	numRows := len(mat)
	numCols := len(mat[0])

	numToIndex := make([]int, len(arr)+1)
	for i, num := range arr {
		numToIndex[num] = i
	}

	result := math.MaxInt32

	for row := 0; row < numRows; row++ {
		lastElemIdx := math.MinInt32
		for col := 0; col < numCols; col++ {
			idxVal := numToIndex[mat[row][col]]
			lastElemIdx = max(lastElemIdx, idxVal)
		}
		result = min(result, lastElemIdx)
	}

	for col := 0; col < numCols; col++ {
		lastElemIdx := math.MinInt32
		for row := 0; row < numRows; row++ {
			idxVal := numToIndex[mat[row][col]]
			lastElemIdx = max(lastElemIdx, idxVal)
		}
		result = min(result, lastElemIdx)
	}

	return result
}
