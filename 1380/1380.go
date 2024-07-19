// 1380. Lucky Numbers in a Matrix
package lucky_numbers

import (
	"math"
)

// luckyNumbers returns all lucky numbers in the matrix.
// A lucky number is an element of the matrix such that it is the minimum element in its row and maximum in its column.
func luckyNumbers(matrix [][]int) []int {
	luckyNums := make([]int, 0)

outer:
	for _, row := range matrix {
		rowMin := math.MaxInt
		rowMinIdx := 0

		for j, n := range row {
			if n < rowMin {
				rowMin = n
				rowMinIdx = j
			}
		}

		for i := 0; i < len(matrix); i++ {
			if matrix[i][rowMinIdx] > rowMin {
				continue outer
			}
		}

		luckyNums = append(luckyNums, rowMin)
	}

	return luckyNums
}
