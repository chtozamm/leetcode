// 2976. Minimum Cost to Convert String I
package minimum_cost

import (
	"container/heap"
)

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	graph := make(map[byte]map[byte]int)
	for i := range original {
		from, to := original[i], changed[i]
		if graph[from] == nil {
			graph[from] = make(map[byte]int)
		}
		if existingCost, exists := graph[from][to]; !exists || cost[i] < existingCost {
			graph[from][to] = cost[i]
		}
	}

	cache := make(map[byte]map[byte]int)

	totalCost := 0
	for i := range source {
		if source[i] == target[i] {
			continue
		}
		if cachedCost, exists := getCache(cache, source[i], target[i]); exists {
			totalCost += cachedCost
			continue
		}

		changeCost := dijkstra(graph, source[i], target[i], cache)
		if changeCost == -1 {
			return -1
		}

		totalCost += changeCost
	}

	return int64(totalCost)
}

func dijkstra(graph map[byte]map[byte]int, source, target byte, cache map[byte]map[byte]int) int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Edge{source, 0})
	visited := make(map[byte]bool)

	for pq.Len() > 0 {
		root := heap.Pop(pq).(Edge)
		if visited[root.char] {
			continue
		}
		visited[root.char] = true
		setCache(cache, source, root.char, root.cost)

		if root.char == target {
			return root.cost
		}

		for neighbor, edgeCost := range graph[root.char] {
			if !visited[neighbor] {
				newCost := root.cost + edgeCost
				heap.Push(pq, Edge{neighbor, newCost})
			}
		}
	}

	return -1
}

func getCache(cache map[byte]map[byte]int, from, to byte) (int, bool) {
	if _, exists := cache[from]; exists {
		if cost, exists := cache[from][to]; exists {
			return cost, true
		}
	}
	return 0, false
}

func setCache(cache map[byte]map[byte]int, from, to byte, cost int) {
	if cache[from] == nil {
		cache[from] = make(map[byte]int)
	}
	cache[from][to] = cost
}

type Edge struct {
	char byte
	cost int
}

type PriorityQueue []Edge

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Edge))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
