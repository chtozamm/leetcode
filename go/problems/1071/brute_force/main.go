package main

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
	len1, len2 := len(str1), len(str2)

	var valid func(k int) bool
	valid = func(k int) bool {
		if len1%k > 0 || len2%k > 0 {
			return false
		}
		base := str1[:k]
		return strings.ReplaceAll(str1, base, "") == "" && strings.ReplaceAll(str2, base, "") == ""
	}

	for i := min(len1, len2); i >= 1; i-- {
		if valid(i) {
			return str1[:i]
		}
	}

	return ""
}
