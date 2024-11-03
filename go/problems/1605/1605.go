// 1605. Find Valid Matrix Given Row and Column Sums
package valid_matrix

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	m := len(rowSum)
	n := len(colSum)

	currRowSum := make([]int, m)
	currColSum := make([]int, n)
	restoredMatrix := make([][]int, m)

	for i := 0; i < m; i++ {
		restoredMatrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			restoredMatrix[i][j] = min(rowSum[i]-currRowSum[i], colSum[j]-currColSum[j])
			currRowSum[i] += restoredMatrix[i][j]
			currColSum[j] += restoredMatrix[i][j]
		}
	}

	return restoredMatrix
}
