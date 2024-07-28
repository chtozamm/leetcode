// 2045. Second Minimum Time to Reach Destination
package minimum_time

import (
	"container/heap"
)

func secondMinimumDijkstra(n int, edges [][]int, time int, change int) int {
	// Create adjacency list to represent the graph
	adj := make([][]int, n+1)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		adj[from] = append(adj[from], to)
		adj[to] = append(adj[to], from)
	}

	// Create slices to store the first and second shortest times
	shortestTime1 := make([]int, n+1)
	shortestTime2 := make([]int, n+1)
	for i := 1; i <= n; i++ {
		shortestTime1[i] = -1
		shortestTime2[i] = -1
	}

	// Initialize the priority queue and start with node 1
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: 1, timeTaken: 0})
	shortestTime1[1] = 0

	// Process the queue until empty
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item)
		node := curr.node
		timeTaken := curr.timeTaken

		// Calculate the wait time if the signal is red
		if (timeTaken/change)%2 == 1 {
			timeTaken = change*(timeTaken/change+1) + time
		} else {
			timeTaken += time
		}

		// Explore neighbors
		for _, neighbor := range adj[node] {
			if shortestTime1[neighbor] == -1 {
				// Update the first minimum time for the neighbor
				shortestTime1[neighbor] = timeTaken
				heap.Push(pq, &Item{node: neighbor, timeTaken: timeTaken})
			} else if shortestTime2[neighbor] == -1 && shortestTime1[neighbor] != timeTaken {
				// Update the second minimum time for the neighbor
				if neighbor == n {
					return timeTaken
				}
				shortestTime2[neighbor] = timeTaken
				heap.Push(pq, &Item{node: neighbor, timeTaken: timeTaken})
			}
		}
	}

	return 0
}

type Item struct {
	node, timeTaken int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].timeTaken < pq[j].timeTaken
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
