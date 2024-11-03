// 2392. Build a Matrix With Conditions
package matrix

type NodeStatus int

const (
	Unvisited NodeStatus = iota
	Visiting
	Visited
)

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	orderRows := topoSort(rowConditions, k)
	orderColumns := topoSort(colConditions, k)

	if len(orderRows) == 0 || len(orderColumns) == 0 {
		return [][]int{}
	}

	matrix := make([][]int, k)
	for i := range matrix {
		matrix[i] = make([]int, k)
	}

	posRow := make(map[int]int)
	posCol := make(map[int]int)

	for i, num := range orderRows {
		posRow[num] = i
	}
	for i, num := range orderColumns {
		posCol[num] = i
	}

	for num := 1; num <= k; num++ {
		if row, ok1 := posRow[num]; ok1 {
			if col, ok2 := posCol[num]; ok2 {
				matrix[row][col] = num
			}
		}
	}

	return matrix
}

func topoSort(edges [][]int, n int) []int {
	adj := make(map[int][]int)
	order := []int{}
	visited := make([]NodeStatus, n+1)
	hasCycle := false

	for _, edge := range edges {
		x, y := edge[0], edge[1]
		adj[x] = append(adj[x], y)
	}

	for i := 1; i <= n; i++ {
		if visited[i] == Unvisited {
			dfs(i, adj, visited, &order, &hasCycle)
			if hasCycle {
				return []int{}
			}
		}
	}

	reverse(order)
	return order
}

func dfs(node int, adj map[int][]int, visited []NodeStatus, order *[]int, hasCycle *bool) {
	visited[node] = Visiting
	for _, neighbor := range adj[node] {
		if visited[neighbor] == Unvisited {
			dfs(neighbor, adj, visited, order, hasCycle)
			if *hasCycle {
				return
			}
		} else if visited[neighbor] == Visiting {
			*hasCycle = true
			return
		}
	}
	visited[node] = Visited
	*order = append(*order, node)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
