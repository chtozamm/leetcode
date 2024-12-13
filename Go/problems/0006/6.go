// 6. Zigzag Conversion
package zigzag_pattern

import "strings"

func convert(s string, numRows int) string {
	if numRows <= 1 || numRows >= len(s) {
		return s
	}

	// Create a slice of strings for each row
	rows := make([]strings.Builder, numRows)
	currentRow := 0
	goingDown := false

	for _, char := range s {
		rows[currentRow].WriteRune(char)

		// Change direction if we are at the top or bottom row
		if currentRow == 0 || currentRow == numRows-1 {
			goingDown = !goingDown
		}

		// Move to the next row
		if goingDown {
			currentRow++
		} else {
			currentRow--
		}
	}

	// Combine all rows into a single string
	var result strings.Builder
	for _, row := range rows {
		result.WriteString(row.String())
	}

	return result.String()
}
