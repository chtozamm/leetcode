// 726. Number of Atoms
package atoms

import (
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var index int

func countOfAtoms(formula string) string {
	index = 0 // Reset index to 0 at the beginning of each call (for tests)

	finalMap := parseFormula(formula)

	keys := make([]string, 0, len(finalMap))
	for k := range finalMap {
		keys = append(keys, k)
	}

	// Sort the keys
	sort.Strings(keys)

	// Generate the answer string
	var ans strings.Builder
	for _, atom := range keys {
		ans.WriteString(atom)
		if finalMap[atom] > 1 {
			ans.WriteString(strconv.Itoa(finalMap[atom]))
		}
	}

	return ans.String()
}

func parseFormula(formula string) map[string]int {
	currMap := make(map[string]int)
	length := len(formula)

	for index < length && formula[index] != ')' {
		if formula[index] == '(' {
			index++
			nestedMap := parseFormula(formula)
			for atom, count := range nestedMap {
				currMap[atom] += count
			}
		} else {
			currAtom := parseAtom(formula)
			currCount := parseCount(formula)
			if currCount == 0 {
				currCount = 1
			}
			currMap[currAtom] += currCount
		}
	}

	index++ // to skip ')'
	multiplier := parseCount(formula)
	if multiplier > 0 {
		for atom := range currMap {
			currMap[atom] *= multiplier
		}
	}

	return currMap
}

func parseAtom(formula string) string {
	start := index
	index++
	length := len(formula)
	for index < length && unicode.IsLower(rune(formula[index])) {
		index++
	}
	return formula[start:index]
}

func parseCount(formula string) int {
	start := index
	length := len(formula)
	for index < length && unicode.IsDigit(rune(formula[index])) {
		index++
	}
	if start == index {
		return 0
	}
	count, _ := strconv.Atoi(formula[start:index])
	return count
}
