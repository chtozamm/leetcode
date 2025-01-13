package main

import "unicode"

func reverseVowels(s string) string {
	n := len(s)
	res := []rune(s)

	i, j := 0, n-1
	for i < j {
		for i < j && !isVowel(res[i]) {
			i++
		}
		for i < j && !isVowel(res[j]) {
			j--
		}
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return string(res)
}

func isVowel(char rune) bool {
	char = unicode.ToLower(char)
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}
