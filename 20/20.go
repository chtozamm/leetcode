// 20. Valid Parentheses
package valid_parenthesis

func isValid(s string) bool {
	stack := []rune{}

	parenthesesPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, r := range s {
		switch r {
		case '(', '{', '[':
			stack = append(stack, r)
		case ')', '}', ']':
			// If the stack is empty or the top of the stack does not match the corresponding opening parenthesis, return false
			if len(stack) == 0 || stack[len(stack)-1] != parenthesesPairs[r] {
				return false
			}
			if len(stack) > 0 && stack[len(stack)-1] == parenthesesPairs[r] {
				// Pop the stack
				stack = stack[:len(stack)-1]
			}
		}
	}

	// If the stack is empty, all parentheses were matched; otherwise, return false
	return len(stack) == 0
}
