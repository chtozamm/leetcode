package main

import (
	"container/list"
	"math"
)

func minCost(grid [][]int) int {
	numRows := len(grid)
	numCols := len(grid[0])

	// Initialize the minimum cost grid with maximum integer values
	minCostGrid := make([][]int, numRows)
	for i := range minCostGrid {
		minCostGrid[i] = make([]int, numCols)
		for j := range minCostGrid[i] {
			minCostGrid[i][j] = math.MaxInt32
		}
	}
	minCostGrid[0][0] = 0

	directions := map[int][2]int{
		1: {0, 1},  // Right
		2: {0, -1}, // Left
		3: {1, 0},  // Down
		4: {-1, 0}, // Up
	}

	// Initialize the queue for 0-1 BFS — add zero-cost moves to the front and cost=1 moves to the back
	queue := list.New()
	queue.PushFront([2]int{0, 0})

	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)
		item := elem.Value.([2]int)
		row, col := item[0], item[1]

		for dirIdx, delta := range directions {
			newRow := row + delta[0]
			newCol := col + delta[1]

			cost := 0
			if grid[row][col] != dirIdx {
				cost = 1
			}

			// Check if the new position is within bounds
			if newRow < 0 || newRow >= numRows || newCol < 0 || newCol >= numCols {
				continue
			}

			// If a shorter path is found, update the minimum cost for the new position and add it to the queue
			if minCostGrid[row][col]+cost < minCostGrid[newRow][newCol] {
				minCostGrid[newRow][newCol] = minCostGrid[row][col] + cost

				if cost == 0 {
					queue.PushFront([2]int{newRow, newCol})
				} else {
					queue.PushBack([2]int{newRow, newCol})
				}
			}
		}
	}

	return minCostGrid[numRows-1][numCols-1]
}
