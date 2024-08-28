// 125. Valid Palindrome
package palindrome

import "unicode"

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		for start < end && !unicode.IsLetter(rune(s[start])) && !unicode.IsDigit(rune(s[start])) {
			start++
		}
		for start < end && !unicode.IsLetter(rune(s[end])) && !unicode.IsDigit(rune(s[end])) {
			end--
		}
		if unicode.ToLower(rune(s[start])) != unicode.ToLower(rune(s[end])) {
			return false
		}
		start++
		end--
	}

	return true
}
