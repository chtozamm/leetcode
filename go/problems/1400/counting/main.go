package main

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

	if len(s) == k {
		return true
	}

	freq := [26]int{}
	for _, char := range s {
		freq[char-'a']++
	}

	oddCount := 0
	for _, count := range freq {
		if count%2 == 1 {
			oddCount++
		}
	}

	return oddCount <= k
}
