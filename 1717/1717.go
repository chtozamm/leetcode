// 1717. Maximum Score From Removing Substrings
package gain

func maximumGain(s string, x int, y int) int {
	var firstSubstr, secondSubstr string
	var firstScore, secondScore int
	if x > y {
		firstSubstr, secondSubstr = "ab", "ba"
		firstScore, secondScore = x, y
	} else {
		firstSubstr, secondSubstr = "ba", "ab"
		firstScore, secondScore = y, x
	}

	score := 0
	stack := []byte{}
	first, second := firstSubstr[0], firstSubstr[1]

	// First pass: remove all instances of firstSubstr and accumulate the score
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == first && s[i] == second {
			stack = stack[:len(stack)-1]
			score += firstScore
		} else {
			stack = append(stack, s[i])
		}
	}

	// Prepare for the second pass
	s = string(stack)
	stack = stack[:0]
	first, second = secondSubstr[0], secondSubstr[1]

	// Second pass: remove all instances of secondSubstr and accumulate the score
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == first && s[i] == second {
			stack = stack[:len(stack)-1]
			score += secondScore
		} else {
			stack = append(stack, s[i])
		}
	}

	return score
}
