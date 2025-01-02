package main

func vowelStrings(words []string, queries [][]int) []int {
	prefixSum := make([]int, len(words))
	sum := 0

	for i, word := range words {
		if isVowel(word[0]) && isVowel(word[len(word)-1]) {
			sum++
		}
		prefixSum[i] = sum
	}

	output := make([]int, len(queries))

	for i, query := range queries {
		left, right := query[0], query[1]
		if left > 0 {
			output[i] = prefixSum[right] - prefixSum[left-1]
		} else {
			output[i] = prefixSum[right]
		}
	}

	return output
}

func isVowel(char byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	for _, vowel := range vowels {
		if char == vowel {
			return true
		}
	}
	return false
}
