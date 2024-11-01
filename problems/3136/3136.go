// 3136. Valid Word
package valid_word

import (
	"regexp"
	"unicode"
)

func isValid(word string) bool {
	// Check if the string contains only digits and English letters and has at least 3 characters
	matched, _ := regexp.MatchString("^[A-Za-z0-9]{3,}$", word)
	if !matched {
		return false
	}

	// Check for at least one vowel and one consonant
	var hasVowel, hasConsonant bool
	for _, char := range word {
		if unicode.IsLetter(char) {
			switch unicode.ToLower(char) {
			case 'a', 'e', 'i', 'o', 'u':
				hasVowel = true
			default:
				hasConsonant = true
			}
		}
	}
	return hasVowel && hasConsonant
}
