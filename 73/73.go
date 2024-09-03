// 73. Set Matrix Zeroes
package matrix_zeroes

func setZeroes(matrix [][]int) {
	firstRowHasZeroes := false
	firstColumnHasZeroes := false

	// Check for zeroes in the first row
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			firstRowHasZeroes = true
			break
		}
	}

	// Check for zeroes in the first column
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstColumnHasZeroes = true
			break
		}
	}

	// Check for zeroes in the rest of the matrix
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// Nullify columns based on the first row
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}

	// Nullify rows based on the first column
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[i]); j++ {
				matrix[i][j] = 0
			}
		}
	}

	// Nullify first row
	if firstRowHasZeroes {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	// Nullify first column
	if firstColumnHasZeroes {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
