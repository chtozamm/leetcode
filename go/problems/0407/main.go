package main

import "container/heap"

func trapRainWater(heightMap [][]int) int {
	numRows := len(heightMap)
	numCols := len(heightMap[0])

	dRow := [4]int{0, 0, -1, 1}
	dCol := [4]int{-1, 1, 0, 0}

	visited := make(map[[2]int]bool)

	// Priority queue to process boundary cells in increasing height order
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Add the first and last row cells to the boundary and mark them as visited
	for i := 0; i < numCols; i++ {
		heap.Push(pq, &Item{height: heightMap[0][i], row: 0, col: i})
		heap.Push(pq, &Item{height: heightMap[numRows-1][i], row: numRows - 1, col: i})

		visited[[2]int{0, i}] = true
		visited[[2]int{numRows - 1, i}] = true
	}

	// Add the first and last column cells to the boundary and mark them as visited
	for i := 0; i < numRows; i++ {
		heap.Push(pq, &Item{height: heightMap[i][0], row: i, col: 0})
		heap.Push(pq, &Item{height: heightMap[i][numCols-1], row: i, col: numCols - 1})

		visited[[2]int{i, 0}] = true
		visited[[2]int{i, numCols - 1}] = true
	}

	totalWaterVolume := 0

	for pq.Len() > 0 {
		currentCell := heap.Pop(pq).(*Item)

		currentRow := currentCell.row
		currentCol := currentCell.col
		minBoundaryHeight := currentCell.height

		for dir := 0; dir < 4; dir++ {
			neighborRow := currentRow + dRow[dir]
			neighborCol := currentCol + dCol[dir]

			if !isWithinBounds(neighborRow, neighborCol, numRows, numCols) {
				continue
			}

			if visited[[2]int{neighborRow, neighborCol}] {
				continue
			}

			neighborHeight := heightMap[neighborRow][neighborCol]
			if neighborHeight < minBoundaryHeight {
				totalWaterVolume += minBoundaryHeight - neighborHeight
			}

			heap.Push(pq, &Item{
				height: max(neighborHeight, minBoundaryHeight),
				row:    neighborRow,
				col:    neighborCol,
			})
			visited[[2]int{neighborRow, neighborCol}] = true
		}
	}

	return totalWaterVolume
}

func isWithinBounds(row, col, numRows, numCols int) bool {
	return row >= 0 && col >= 0 && row < numRows && col < numCols
}
