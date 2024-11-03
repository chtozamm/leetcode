// 2045. Second Minimum Time to Reach Destination
package minimum_time

import (
	"container/list"
)

func secondMinimumBFS(n int, edges [][]int, time int, change int) int {
	// Create adjacency list to represent the graph
	adj := make(map[int][]int)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		adj[from] = append(adj[from], to)
		adj[to] = append(adj[to], from)
	}

	// Initialize distances for the shortest and second shortest paths
	dist1 := make([]int, n+1)
	dist2 := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist1[i] = -1
		dist2[i] = -1
	}

	// Use a queue for BFS, starting from node 1 with frequency 1 (first visit)
	queue := list.New()
	queue.PushBack([2]int{1, 1})
	dist1[1] = 0

	for queue.Len() > 0 {
		temp := queue.Remove(queue.Front()).([2]int)
		node := temp[0]
		freq := temp[1]

		// Determine the time taken to reach the current node
		var timeTaken int
		if freq == 1 {
			timeTaken = dist1[node]
		} else {
			timeTaken = dist2[node]
		}

		// Calculate the wait time if the signal is red
		if (timeTaken/change)%2 == 1 {
			timeTaken = change*(timeTaken/change+1) + time
		} else {
			timeTaken += time
		}

		// Explore neighbors
		if neighbors, ok := adj[node]; ok {
			for _, neighbor := range neighbors {
				if dist1[neighbor] == -1 {
					// If the shortest path to the neighbor has not been found
					dist1[neighbor] = timeTaken
					queue.PushBack([2]int{neighbor, 1})
				} else if dist2[neighbor] == -1 && dist1[neighbor] != timeTaken {
					// If the second shortest path to the neighbor has not been found and it's different from the first
					if neighbor == n {
						return timeTaken
					}
					dist2[neighbor] = timeTaken
					queue.PushBack([2]int{neighbor, 2})
				}
			}
		}
	}

	return 0
}
