package main

func maxScore(s string) int {
	var best, left, right int

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			right++
		}
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			left++
		} else {
			right--
		}
		best = max(best, left+right)
	}

	return best
}
