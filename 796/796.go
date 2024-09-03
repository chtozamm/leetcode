// 796. Rotate String
package rotate_string

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	// Check all rotations
	for i := 0; i < len(s); i++ {
		rotated := s[i:] + s[:i]
		if rotated == goal {
			return true
		}
	}

	return false
}

/* Alternative solution using substring check:
func rotateString(s string, goal string) bool {
  if len(s) != len(goal) {
		return false
	}

	double := s + s
	return strings.Contains(double, goal)
} */
