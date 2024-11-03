// 68. Text Justification
package justification

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var rows []string
	var curr []string
	totalLetters := 0

	for _, word := range words {
		// Check if adding the next word exceeds the maxWidth
		if len(word)+len(curr)+totalLetters > maxWidth {
			rows = append(rows, justify(curr, totalLetters, maxWidth))
			curr = []string{word}
			totalLetters = len(word)
		} else {
			curr = append(curr, word)
			totalLetters += len(word)
		}
	}

	// Handle the last row
	lastRow := strings.Join(curr, " ")
	lastRow += strings.Repeat(" ", maxWidth-len(lastRow))
	rows = append(rows, lastRow)

	return rows
}

func justify(words []string, totalLetters, maxWidth int) string {
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", maxWidth-len(words[0]))
	}

	totalSpaces := maxWidth - totalLetters
	spaceBetweenWords := totalSpaces / (len(words) - 1)
	extraSpaces := totalSpaces % (len(words) - 1)

	var builder strings.Builder
	for i, word := range words {
		builder.WriteString(word)
		if i < len(words)-1 {
			builder.WriteString(strings.Repeat(" ", spaceBetweenWords))
			if i < extraSpaces {
				builder.WriteString(" ")
			}
		}
	}

	return builder.String()
}
