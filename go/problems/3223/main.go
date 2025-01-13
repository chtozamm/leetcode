package main

func minimumLength(s string) int {
	n := len(s)
	if n < 3 {
		return n
	}

	counts := [26]int{}
	for _, char := range s {
		counts[char-'a']++
		if counts[char-'a'] > 2 {
			counts[char-'a'] -= 2
			n -= 2
		}
	}

	return n
}
