package main

func countServers(grid [][]int) int {
	numRows := len(grid)
	numCols := len(grid[0])

	rowCounts := make([]int, numCols)
	colCounts := make([]int, numRows)

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if grid[row][col] == 1 {
				rowCounts[col]++
				colCounts[row]++
			}
		}
	}

	ans := 0

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if grid[row][col] == 1 {
				if rowCounts[col] > 1 || colCounts[row] > 1 {
					ans++
				}
			}
		}
	}

	return ans
}
