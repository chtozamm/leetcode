package atoms

import "testing"

func BenchmarkCountOfAtoms(b *testing.B) {
	testCases := []struct {
		input string
	}{
		{input: "H2O"},
		{input: "Mg(OH)2"},
		{input: "K4(ON(SO3)2)2"},
		{input: "NaCl"},
		{input: "C6H12O6"},
		{input: "Fe(C5H5)2"},
		{input: "U(C6H5CO)3"},
		{input: "B(OH)3"},
	}

	for _, tc := range testCases {
		b.Run("Recursion-"+tc.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				countOfAtomsRecursion(tc.input)
			}
		})
		b.Run("Stack-"+tc.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				countOfAtomsStack(tc.input)
			}
		})
		b.Run("Regex-"+tc.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				countOfAtomsRegex(tc.input)
			}
		})
	}
}
