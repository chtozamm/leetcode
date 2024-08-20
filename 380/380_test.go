package randomized_set

import (
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	rs := Constructor()

	operations := []string{"insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"}
	values := [][]int{{1}, {2}, {2}, {}, {1}, {2}, {}}
	expectedOutput := []interface{}{true, false, true, 1, true, false, 2}
	actualOutput := make([]interface{}, 0)

	for i, op := range operations {
		switch op {
		case "insert":
			val := values[i][0]
			result := rs.Insert(val)
			actualOutput = append(actualOutput, result)
		case "remove":
			val := values[i][0]
			result := rs.Remove(val)
			actualOutput = append(actualOutput, result)
		case "getRandom":
			result := rs.GetRandom()
			actualOutput = append(actualOutput, result)
		}
	}

	for i, expected := range expectedOutput {
		if expected != actualOutput[i] {
			if i == 3 {
				if actualOutput[i] == 1 || actualOutput[i] == 2 {
					continue
				}
			}
			t.Errorf("Expected %v, but got %v at operation %d", expected, actualOutput[i], i)
		}
	}
}
