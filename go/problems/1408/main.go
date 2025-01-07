package main

func stringMatching(words []string) []string {
	ans := make([]string, 0)

	for i := range words {
		lps := computeLPSArray(words[i])

		for j := range words {
			if i == j {
				continue
			}

			if isSubstringOf(words[i], words[j], lps) {
				ans = append(ans, words[i])
				break
			}
		}
	}

	return ans
}
