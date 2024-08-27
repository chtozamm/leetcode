// 28. Find the Index of the First Occurrence in a String
package occurence

import "strings"

func strStr(haystack string, needle string) int {
	for i := range haystack {
		if strings.HasPrefix(haystack[i:], needle) {
			return i
		}
	}

	return -1
}
