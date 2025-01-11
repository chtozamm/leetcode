package main

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	res := strings.Builder{}
	i := 0

	for i < len(word1) && i < len(word2) {
		res.WriteByte(word1[i])
		res.WriteByte(word2[i])
		i++
	}

	if i < len(word1) {
		res.WriteString(word1[i:])
	}

	if i < len(word2) {
		res.WriteString(word2[i:])
	}

	return res.String()
}
