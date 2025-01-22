package main

import "container/list"

func highestPeak(isWater [][]int) [][]int {
	numRows := len(isWater)
	numCols := len(isWater[0])

	heights := make([][]int, numRows)
	queue := list.New()

	for i := 0; i < numRows; i++ {
		heights[i] = make([]int, numCols)
		for j := 0; j < numCols; j++ {
			if isWater[i][j] == 1 {
				heights[i][j] = 0
				queue.PushBack([2]int{i, j})

			} else {
				heights[i][j] = -1
			}
		}
	}

	dRow := [4]int{0, 1, 0, -1}
	dCol := [4]int{1, 0, -1, 0}

	for queue.Len() > 0 {
		// Multi-source BFS
		queueLen := queue.Len()
		for i := 0; i < queueLen; i++ {
			item := queue.Front()
			queue.Remove(item)
			cell := item.Value.([2]int)
			row, col := cell[0], cell[1]

			for d := 0; d < 4; d++ {
				neighborRow := row + dRow[d]
				neighborCol := col + dCol[d]

				if neighborRow >= 0 && neighborRow < numRows &&
					neighborCol >= 0 && neighborCol < numCols &&
					heights[neighborRow][neighborCol] == -1 {
					heights[neighborRow][neighborCol] = heights[row][col] + 1
					queue.PushBack([2]int{neighborRow, neighborCol})
				}
			}
		}
	}

	return heights
}
