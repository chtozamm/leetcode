// 726. Number of Atoms
package atoms

import (
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func countOfAtomsStack(formula string) string {
	type atomCount map[string]int
	stack := []atomCount{make(atomCount)}
	index := 0
	length := len(formula)

	for index < length {
		if formula[index] == '(' {
			stack = append(stack, make(atomCount))
			index++
		} else if formula[index] == ')' {
			currMap := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			index++
			multiplierStr := ""
			for index < length && unicode.IsDigit(rune(formula[index])) {
				multiplierStr += string(formula[index])
				index++
			}
			multiplier := 1
			if multiplierStr != "" {
				multiplier, _ = strconv.Atoi(multiplierStr)
			}
			for atom, count := range currMap {
				stack[len(stack)-1][atom] += count * multiplier
			}
		} else {
			currAtom := string(formula[index])
			index++
			for index < length && unicode.IsLower(rune(formula[index])) {
				currAtom += string(formula[index])
				index++
			}
			currCountStr := ""
			for index < length && unicode.IsDigit(rune(formula[index])) {
				currCountStr += string(formula[index])
				index++
			}
			currCount := 1
			if currCountStr != "" {
				currCount, _ = strconv.Atoi(currCountStr)
			}
			stack[len(stack)-1][currAtom] += currCount
		}
	}

	finalMap := stack[0]
	keys := make([]string, 0, len(finalMap))
	for k := range finalMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var ans strings.Builder
	for _, atom := range keys {
		ans.WriteString(atom)
		if finalMap[atom] > 1 {
			ans.WriteString(strconv.Itoa(finalMap[atom]))
		}
	}

	return ans.String()
}
