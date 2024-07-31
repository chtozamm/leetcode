// 1105. Filling Bookcase Shelves
package bookshelf

import "math"

func minHeightShelves(books [][]int, shelfWidth int) int {
	// Cache to store previous computations
	memo := make([]int, len(books))
	for i := range memo {
		memo[i] = -1
	}

	return helper(books, shelfWidth, memo, 0)
}

func helper(books [][]int, shelfWidth int, memo []int, i int) int {
	if i == len(books) {
		return 0
	}

	// Return answer if already computed
	if memo[i] != -1 {
		return memo[i]
	}

	currentWidth := 0
	maxHeight := 0
	minHeight := math.MaxInt

	// Try placing books[i] to books[j] on the current shelf
	for j := i; j < len(books); j++ {
		currentWidth += books[j][0]
		if currentWidth > shelfWidth {
			break
		}
		maxHeight = max(maxHeight, books[j][1])
		minHeight = min(minHeight, maxHeight+helper(books, shelfWidth, memo, j+1))
	}

	// Store result in cache
	memo[i] = minHeight
	return memo[i]
}
