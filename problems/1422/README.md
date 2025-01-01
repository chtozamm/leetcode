### Benchmark

| Approach                                      | Time Complexity | Space Complexity | Time        | Memory | Allocations |
| --------------------------------------------- | --------------- | ---------------- | ----------- | ------ | ----------- |
| [Brute Force](./brute_force)                  | O(n²)           | O(1)             | 4447 ns/op  | 0 B/op | 0 allocs/op |
| [Count Left Zeros and Right Ones](./improved) | O(n)            | O(1)             | 501.1 ns/op | 0 B/op | 0 allocs/op |
| [One Pass](./one_pass)                        | O(n)            | O(1)             | 222.8 ns/op | 0 B/op | 0 allocs/op |

### Test cases

```go
type testCase struct {
	s        string
	expected int
}

var testCases = []testCase{
	{"011101", 5},
	{"00111", 5},
	{"1111", 3},
	{"0000000000", 9},
	{"1111111111", 9},
	{"0101010101", 6},
	{"1010101010", 5},
	{"000111000111", 9},
	{"111000111000", 6},
	{"01010101010101010101", 11},
	{"11111111110000000000", 9},
	{"00000000001111111111", 20},
	{"0101010101010101010101010101010101", 18},
	{"0000000000000000000000000000000000000000", 39},
}
```
