package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)

	maxCandies := -1
	for i := range candies {
		maxCandies = max(maxCandies, candies[i])
	}

	res := make([]bool, n)
	for i := range candies {
		if candies[i]+extraCandies >= maxCandies {
			res[i] = true
		}
	}

	return res
}
