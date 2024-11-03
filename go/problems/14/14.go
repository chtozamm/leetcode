// 14. Longest Common Prefix
package common_prefix

import "strings"

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for _, str := range strs[1:] {
		// Reduce the prefix until it matches the start of the current string
		for !strings.HasPrefix(str, prefix) {
			prefix = prefix[:len(prefix)-1]
			// If the prefix is reduced to an empty string, return immediately
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}
