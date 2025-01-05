package main

import "strings"

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	diffArray := make([]int, n)

	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]
		if direction == 1 {
			diffArray[start]++
			if end+1 < n {
				diffArray[end+1]--
			}
		} else {
			diffArray[start]--
			if end+1 < n {
				diffArray[end+1]++
			}
		}
	}

	numberOfShifts := 0
	ans := strings.Builder{}

	for i := 0; i < n; i++ {
		numberOfShifts = (numberOfShifts + diffArray[i]) % 26
		if numberOfShifts < 0 {
			numberOfShifts += 26
		}
		newChar := 'a' + (s[i]-'a'+byte(numberOfShifts))%26
		ans.WriteByte(newChar)
	}

	return ans.String()
}
