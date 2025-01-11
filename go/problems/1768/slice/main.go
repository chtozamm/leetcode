package main

func mergeAlternately(word1 string, word2 string) string {
	n := max(len(word1), len(word2))
	res := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		if i < len(word1) {
			res = append(res, word1[i])
		}
		if i < len(word2) {
			res = append(res, word2[i])
		}
	}

	return string(res)
}
