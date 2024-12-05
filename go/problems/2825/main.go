package main

func canMakeSubsequence(str1 string, str2 string) bool {
	cursor := 0

	for i := 0; i < len(str1); i++ {
		if cursor >= len(str2) {
			break
		}

		curr := str1[i]
		next := curr + 1
		goal := str2[cursor]

		if next > 'z' {
			next = 'a'
		}

		if curr == goal || next == goal {
			cursor++
		}
	}

	return cursor == len(str2)
}
