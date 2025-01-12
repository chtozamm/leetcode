package main

func canBeValid(s string, locked string) bool {
	n := len(s)

	if n%2 == 1 {
		return false
	}

	// First pass: left to right
	balance := 0
	for i := range s {
		if locked[i] == '0' || s[i] == '(' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
	}

	// Second pass: right to left
	balance = 0
	for i := n - 1; i >= 0; i-- {
		if locked[i] == '0' || s[i] == ')' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
	}

	return true
}
