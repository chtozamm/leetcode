// 205. Isomorphic Strings
package isomophic_strings

func isIsomorphic(s string, t string) bool {
	charMapST := make(map[rune]rune)
	charMapTS := make(map[rune]rune)

	for i, charS := range s {
		charT := rune(t[i])

		if mappedChar, ok := charMapST[charS]; ok {
			if mappedChar != charT {
				return false
			}
		} else {
			charMapST[charS] = charT
		}

		if mappedChar, ok := charMapTS[charT]; ok {
			if mappedChar != charS {
				return false
			}
		} else {
			charMapTS[charT] = charS
		}
	}

	return true
}
