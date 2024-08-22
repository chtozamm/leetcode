// 135. Candy
package candy

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1 // Each child gets at least one candy
	}

	// First pass: left to right
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Second pass: right to left
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candies[i] <= candies[i+1] {
				candies[i] = candies[i+1] + 1
			}
		}
	}

	// Calculate total candies
	totalCandies := 0
	for _, candy := range candies {
		totalCandies += candy
	}

	return totalCandies
}
