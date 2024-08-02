// 13. Roman to Integer
package roman_to_int

func romanToInt(s string) int {
	romToIntMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	num := 0
	n := len(s)

	for i := 0; i < n; i++ {
		currentValue := romToIntMap[s[i]]
		if i+1 < n && currentValue < romToIntMap[s[i+1]] {
			num -= currentValue
		} else {
			num += currentValue
		}
	}

	return num
}
