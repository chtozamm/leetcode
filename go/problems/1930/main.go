package main

func countPalindromicSubsequence(s string) int {
	first := [26]int{}
	last := [26]int{}

	for i := 0; i < 26; i++ {
		first[i] = -1
		last[i] = -1
	}

	for i := 0; i < len(s); i++ {
		idx := int(s[i]) - int('a')
		if first[idx] == -1 {
			first[idx] = i
		}
		last[idx] = i
	}

	ans := 0

	for i := 0; i < 26; i++ {
		if first[i] == -1 {
			continue
		}

		between := make(map[byte]interface{})

		for j := first[i] + 1; j < last[i]; j++ {
			between[s[j]] = struct{}{}
		}

		ans += len(between)
	}

	return ans
}
