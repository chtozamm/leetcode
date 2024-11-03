// 66. Plus One
package plus_one

func plusOne(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	// If all digits were 9, create a new slice with an extra space
	result := make([]int, n+1)
	result[0] = 1
	return result
}
