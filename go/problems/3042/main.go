package main

func countPrefixSuffixPairs(words []string) int {
	ans := 0

	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				ans++
			}
		}
	}

	return ans
}

func isPrefixAndSuffix(sub, str string) bool {
	if len(str) < len(sub) {
		return false
	}

	// Check prefix
	for i := range sub {
		if sub[i] != str[i] {
			return false
		}
	}

	// Check suffix
	for i := range sub {
		if sub[len(sub)-1-i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}
