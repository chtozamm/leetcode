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
		got := countOfAtoms(tc.input)
		if got != tc.expected {
			t.Errorf("got %q; expected %q", got, tc.expected)
		}
	}
}
