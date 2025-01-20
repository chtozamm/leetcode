package main

func firstCompleteIndex(arr []int, mat [][]int) int {
	numRows := len(mat)
	numCols := len(mat[0])

	rowCount := make([]int, numRows)
	colCount := make([]int, numCols)

	numToPos := make([][2]int, len(arr)+1)
	for row := range mat {
		for col := range mat[row] {
			numToPos[mat[row][col]] = [2]int{row, col}
		}
	}

	for i, num := range arr {
		pos := numToPos[num]
		row := pos[0]
		col := pos[1]

		rowCount[row]++
		colCount[col]++

		if rowCount[row] == numCols || colCount[col] == numRows {
			return i
		}
	}

	return -1
}
