// 20. Valid Parentheses
package valid_parenthesis

func isValid(s string) bool {
	stack := []rune{}

	parenthesesPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		default:
			if len(stack) == 0 || stack[len(stack)-1] != parenthesesPairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
