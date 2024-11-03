// 1190. Reverse Substrings Between Each Pair of Parentheses
package parentheses

import (
	"strings"

	"github.com/chtozamm/leetcode/utils"
)

func reverseParentheses(s string) string {
	stack := []string{} // used to store states of string builder
	var current strings.Builder

	for _, char := range s {
		if char == '(' {
			// Push the current string to the stack and reset the string builder
			stack = append(stack, current.String())
			current.Reset()
		} else if char == ')' {
			// Reverse the current builder content and append to the last string in the stack
			reversed := utils.ReverseString(current.String())
			// Pop the last element from the stack and append the reversed string to it
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			current.Reset()
			current.WriteString(last)
			current.WriteString(reversed)
		} else {
			current.WriteRune(char)
		}
	}

	return current.String()
}
