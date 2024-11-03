// 1823. Find the Winner of the Circular Game
package circular_game

type node struct {
	index int
	next  *node
}

func findTheWinner(n int, k int) int {
	// Edge case: players are removed 1 by 1
	if k == 1 {
		return n
	}

	playersIdxs := make([]int, n)
	for i := 1; i <= n; i++ {
		playersIdxs[i-1] = i
	}

	head := createCircularLinkedList(playersIdxs)
	current := head
	var prev *node

	for current.index != current.next.index {
		for i := 1; i < k; i++ {
			prev = current         // Node previous to the node to delete
			current = current.next // Node to delete
		}
		// Remove node by dereferencing
		prev.next = current.next
		current = current.next
	}

	return current.index
}

func createCircularLinkedList(vals []int) *node {
	if len(vals) == 0 {
		return nil
	}
	head := &node{index: vals[0]}
	current := head
	for i, val := range vals[1:] {
		current.next = &node{index: val}
		current = current.next
		// If it's a tail, point it to head
		if i == len(vals)-2 {
			current.next = head
		}
	}
	return head
}
