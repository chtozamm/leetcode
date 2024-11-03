// 383. Ransom Note
package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	lettersCount := make(map[rune]int)
	for _, letter := range magazine {
		lettersCount[letter]++
	}

	for _, letter := range ransomNote {
		if lettersCount[letter] == 0 {
			return false
		}
		lettersCount[letter]--
	}

	return true
}
