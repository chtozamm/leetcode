// 58. Length of Last Word
package last_word

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	lastSpaceIndex := strings.LastIndex(s, " ")
	if lastSpaceIndex == -1 {
		return len(s)
	}

	return len(s) - lastSpaceIndex - 1
}
