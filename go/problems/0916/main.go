package main

func wordSubsets(words1 []string, words2 []string) []string {
	requiredChars := [26]int{}
	for _, word := range words2 {
		for i, count := range countChars(word) {
			requiredChars[i] = max(requiredChars[i], count)
		}
	}

	ans := make([]string, 0)

outer:
	for _, word := range words1 {
		wordChars := countChars(word)

		for i, count := range wordChars {
			if count < requiredChars[i] {
				continue outer
			}
		}

		ans = append(ans, word)
	}

	return ans
}

func countChars(word string) [26]int {
	chars := [26]int{}
	for _, char := range word {
		chars[char-'a']++
	}
	return chars
}
