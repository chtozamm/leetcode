// 1957. Delete Characters to Make Fancy String
package fancy_string

import "strings"

func makeFancyString(s string) string {
	prev := s[0]
	count := 1

	var fancyString strings.Builder
	fancyString.WriteByte(prev)

	for _, curr := range s[1:] {
		if curr == rune(prev) {
			count++

			if count > 2 {
				continue
			}
		} else {
			prev = byte(curr)
			count = 1
		}

		fancyString.WriteRune(curr)
	}

	return fancyString.String()
}
