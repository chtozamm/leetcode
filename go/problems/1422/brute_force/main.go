package main

func maxScore(s string) int {
	best := 0

	for i := 0; i < len(s)-1; i++ {
		currScore := 0

		for j := 0; j <= i; j++ {
			if s[j] == '0' {
				currScore++
			}
		}

		for j := i + 1; j < len(s); j++ {
			if s[j] == '1' {
				currScore++
			}
		}

		best = max(best, currScore)
	}

	return best
}
