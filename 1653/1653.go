// 1653. Minimum Deletions to Make String Balanced
package balanced_string

func minimumDeletions(s string) int {
	charStack := []rune{}
	deleteCount := 0

	for _, char := range s {
		// If stack is not empty, top of stack is 'b',
		// and current char is 'a'
		if len(charStack) > 0 && charStack[len(charStack)-1] == 'b' && char == 'a' {
			charStack = charStack[:len(charStack)-1]
			deleteCount++
		} else {
			charStack = append(charStack, char)
		}
	}

	return deleteCount
}
