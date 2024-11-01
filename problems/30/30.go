// 30. Substring with Concatenation of All Words
package substring

func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	wordCount := len(words)
	result := []int{}

	wordCountMap := make(map[string]int)
	for _, word := range words {
		wordCountMap[word]++
	}

	// Iterate over each possible starting point
	for i := 0; i < wordLen; i++ {
		left := i
		count := 0
		seen := make(map[string]int)

		for right := i; right <= len(s)-wordLen; right += wordLen {
			word := s[right : right+wordLen]

			if _, ok := wordCountMap[word]; ok {
				seen[word]++
				count++

				// If the word count exceeds the expected count, move the left pointer
				for seen[word] > wordCountMap[word] {
					leftWord := s[left : left+wordLen]
					seen[leftWord]--
					if seen[leftWord] == 0 {
						delete(seen, leftWord)
					}
					left += wordLen
					count--
				}

				// If all words are found, add the starting index to the result
				if count == wordCount {
					result = append(result, left)
				}
			} else {
				// Reset the window if the word is not in the wordCountMap
				seen = make(map[string]int)
				count = 0
				left = right + wordLen
			}
		}
	}

	return result
}
