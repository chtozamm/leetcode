package atoms

import (
	"testing"
)

type testCase struct {
	input    string
	expected string
}

func TestCountOfAtoms(t *testing.T) {
	testCases := []testCase{
		{input: "H2O", expected: "H2O"},
		{input: "Mg(OH)2", expected: "H2MgO2"},
		{input: "K4(ON(SO3)2)2", expected: "K4N2O14S4"},
	}

	for _, tc := range testCases {
		t.Run(tc.input+" Recursion", func(t *testing.T) {
			got := countOfAtomsRecursion(tc.input)
			if got != tc.expected {
				t.Errorf("Recursion-Based approach: got %q, expected %q", got, tc.expected)
			}
		})

		t.Run(tc.input+" Stack", func(t *testing.T) {
			got := countOfAtomsStack(tc.input)
			if got != tc.expected {
				t.Errorf("Stack-Based approach: got %q, expected %q", got, tc.expected)
			}
		})

		t.Run(tc.input+" Regular Expression", func(t *testing.T) {
			got := countOfAtomsRegex(tc.input)
			if got != tc.expected {
				t.Errorf("Regular Expression-Based approach: got %q, expected %q", got, tc.expected)
			}
		})
	}
}
