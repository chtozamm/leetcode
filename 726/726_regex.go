// 726. Number of Atoms
package atoms

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func countOfAtomsRegex(formula string) string {
	type atomCount map[string]int
	stack := []atomCount{make(atomCount)}

	// Regular expression to extract atom, count, (, ), multiplier
	re := regexp.MustCompile(`([A-Z][a-z]*)(\d*)|(\()|(\))(\d*)`)
	matches := re.FindAllStringSubmatch(formula, -1)

	// Parse the formula
	for _, match := range matches {
		atom := match[1]
		count := match[2]
		leftParen := match[3]
		rightParen := match[4]
		multiplier := match[5]

		if atom != "" {
			num := 1
			if count != "" {
				num, _ = strconv.Atoi(count)
			}
			stack[len(stack)-1][atom] += num
		} else if leftParen != "" {
			stack = append(stack, make(atomCount))
		} else if rightParen != "" {
			currMap := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			mult := 1
			if multiplier != "" {
				mult, _ = strconv.Atoi(multiplier)
			}
			for a, cnt := range currMap {
				stack[len(stack)-1][a] += cnt * mult
			}
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
