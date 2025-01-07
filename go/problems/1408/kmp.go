package main

func computeLPSArray(sub string) []int {
	lps := make([]int, len(sub))
	currIdx := 1
	length := 0

	for currIdx < len(sub) {
		if sub[currIdx] == sub[length] {
			length++
			lps[currIdx] = length
			currIdx++
		} else {
			if length > 0 {
				length = lps[length-1]
			} else {
				currIdx++
			}
		}
	}

	return lps
}

func isSubstringOf(sub string, str string, lps []int) bool {
	strIdx := 0
	subIdx := 0

	for strIdx < len(str) {
		if str[strIdx] == sub[subIdx] {
			strIdx++
			subIdx++
			if subIdx == len(sub) {
				return true
			}
		} else {
			if subIdx > 0 {
				subIdx = lps[subIdx-1]
			} else {
				strIdx++
			}
		}
	}

	return false
}
