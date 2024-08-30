// 3. Longest Substring Without Repeating Characters
package longest_substring

import "strings"

func lengthOfLongestSubstring(s string) int {
	charsSlice := strings.Split(s, "")
	maxCount := 0

	for i := 0; i < len(charsSlice); i++ {
		charsMap := map[string]bool{}
		count := 0
		for j := i; j < len(charsSlice); j++ {
			// Break if the character is present in the map
			// Otherwise add it to the map and increment the count
			if charsMap[charsSlice[j]] {
				break
			} else {
				charsMap[charsSlice[j]] = true
				count++
			}
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
