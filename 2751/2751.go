package robot_collisions

import (
	"slices"
)

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	hpMap := make(map[int]int)
	dirMap := make(map[int]string)

	for i, pos := range positions {
		hpMap[pos] = healths[i]
		dirMap[pos] = string(directions[i])
	}

	// Sort robots by positions
	sortedPositions := make([]int, len(positions))
	copy(sortedPositions, positions)
	slices.Sort(sortedPositions)

	// Stack to keep track of robots
	stack := make([]int, 0, len(positions))

	for _, cur := range sortedPositions {
		for len(stack) > 0 {
			prev := stack[len(stack)-1]

			if dirMap[cur] == "L" && dirMap[prev] == "R" {
				// Collide
				if hpMap[prev] < hpMap[cur] {
					stack = stack[:len(stack)-1] // Remove previous robot
					hpMap[cur]--                 // Decrease current robot's health
					hpMap[prev] = 0              // Previous robot is destroyed
				} else if hpMap[prev] > hpMap[cur] {
					hpMap[prev]--  // Decrease previous robot's health
					hpMap[cur] = 0 // Current robot is destroyed
					break
				} else {
					// Both robots destroy each other
					stack = stack[:len(stack)-1]
					hpMap[prev] = 0
					hpMap[cur] = 0
					break
				}
			} else {
				break
			}
		}

		if hpMap[cur] > 0 {
			stack = append(stack, cur)
		}
	}

	// Restore the order of robots' positions
	finalHps := make([]int, 0, len(positions))
	for _, pos := range positions {
		if hpMap[pos] > 0 {
			finalHps = append(finalHps, hpMap[pos])
		}
	}

	return finalHps
}
