package main

func maxScore(s string) int {
	var zeroes, ones int
	best := -1 << 31

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '1' {
			ones++
		} else {
			zeroes++
		}

		best = max(best, zeroes-ones)
	}

	if s[len(s)-1] == '1' {
		ones++
	}

	return best + ones
}
