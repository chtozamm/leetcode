// 22. Generate Parentheses
package parentheses

func generateParenthesis(n int) []string {
	parentheses := make([]string, 0, 2*n)

	backtrack(&parentheses, "", 0, 0, n)

	return parentheses
}

func backtrack(parentheses *[]string, current string, openCount, closeCount, n int) {
	// Base Case: Check if the current string has reached the maximum length
	if len(current) == 2*n {
		*parentheses = append(*parentheses, current)
		return
	}

	// Recursive Case: Add '(' or ')' to the current string and backtrack
	if openCount < n {
		backtrack(parentheses, current+"(", openCount+1, closeCount, n)
	}

	if closeCount < openCount {
		backtrack(parentheses, current+")", openCount, closeCount+1, n)
	}
}
