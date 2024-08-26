// 151. Reverse Words in a String
package reverse_words

import (
	"slices"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	slices.Reverse(words)
	return strings.Join(words, " ")
}
