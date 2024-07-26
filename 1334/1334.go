// 1334. Find the City With the Smallest Number of Neighbors at a Threshold Distance
package city_neighbors

import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// Create matrix to store the shortest distances between all pairs of cities
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
		dist[i][i] = 0
	}

	// Build initial graph
	for _, e := range edges {
		a, b, cost := e[0], e[1], e[2]
		dist[a][b] = cost
		dist[b][a] = cost
	}

	// Floyd-Warshall algorithm
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	minNeighbors := n
	result := -1

	// Find the city with the smallest number of neighbors
	for i := 0; i < n; i++ {
		neighbors := 0
		for j := 0; j < n; j++ {
			if dist[i][j] <= distanceThreshold {
				neighbors++
			}
		}
		if neighbors <= minNeighbors {
			minNeighbors = neighbors
			result = i
		}
	}

	return result
}
