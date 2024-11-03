// 67. Add Binary
package add_binary

func addBinary(a string, b string) string {
	carry := 0
	result := []byte{}

	// Process both strings from the end to the beginning
	for len(a) > 0 || len(b) > 0 || carry > 0 {
		sum := carry
		if len(a) > 0 {
			sum += int(a[len(a)-1] - '0') // Convert char to int
			a = a[:len(a)-1]              // Remove last character
		}
		if len(b) > 0 {
			sum += int(b[len(b)-1] - '0')
			b = b[:len(b)-1]
		}
		// Calculate the new carry and the current bit
		carry = sum / 2
		result = append(result, byte(sum%2+'0'))
	}

	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
