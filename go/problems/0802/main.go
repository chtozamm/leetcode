package main

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	visited := make(map[int]bool)
	isUnsafe := make(map[int]bool)

	for i := 0; i < n; i++ {
		dfsUnsafeNodes(i, graph, visited, isUnsafe)
	}

	safeNodes := make([]int, 0)
	for i := 0; i < n; i++ {
		if !isUnsafe[i] {
			safeNodes = append(safeNodes, i)
		}
	}

	return safeNodes
}

func dfsUnsafeNodes(node int, adjacent [][]int, visited, isUnsafe map[int]bool) bool {
	if isUnsafe[node] {
		return true
	}

	if visited[node] {
		return false
	}

	visited[node] = true
	isUnsafe[node] = true

	for _, neighbor := range adjacent[node] {
		if dfsUnsafeNodes(neighbor, adjacent, visited, isUnsafe) {
			return true
		}
	}

	isUnsafe[node] = false
	return false
}
