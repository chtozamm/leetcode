// 12. Integer to Roman
package int_to_roman

import "strings"

func intToRoman(num int) string {
	sb := strings.Builder{}
	decimals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(decimals); i++ {
		for num >= decimals[i] {
			sb.WriteString(romans[i])
			num -= decimals[i]
		}
	}

	return sb.String()
}
