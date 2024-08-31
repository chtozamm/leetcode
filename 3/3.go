// 3. Longest Substring Without Repeating Characters
package longest_substring

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	left, maxLength := 0, 0

	for right, char := range s {
		// If the character is already in the map and its index is within the current window,
		// move the left pointer to the right
		if index, found := charMap[char]; found && index >= left {
			left = index + 1
		}
		charMap[char] = right
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}
