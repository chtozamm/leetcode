# 726. Number Of Atoms

Problem description could be found [here](https://leetcode.com/problems/number-of-atoms/description/).

**Solutions:**

- [Recursion](./726_recursion.go)
- [Stack](./726_stack.go)
- [Regular Expression](./726_regex.go)

## Benchmark

### For input `H2O`

| Approach   | Time per Operation (ns/op) | Memory Allocated per Operation (B/op) | Allocations per Operation |
|------------|---------------------------:|--------------------------------------:|--------------------------:|
| Recursion  |                         383 |                                  296 |                         4 |
| Stack      |                         443 |                                  312 |                         7 |
| Regex      |                        6711 |                                 6491 |                        47 |

### For input `K4(ON(SO3)2)2`

| Approach   | Time per Operation (ns/op) | Memory Allocated per Operation (B/op) | Allocations per Operation |
|------------|---------------------------:|--------------------------------------:|--------------------------:|
| Recursion  |                        1456 |                                  856 |                         9 |
| Stack      |                        1524 |                                  944 |                        20 |
| Regex      |                       11039 |                                 8470 |                        68 |

## Summary

### Performance

- The recursion-based approach is the fastest and most memory-efficient.
- The stack-based approach is moderately slower than recursion.
- The regex-based approach is significantly slower compared to both recursion and stack.

### Memory Allocations

- Recursion and stack approaches have significantly fewer allocations compared to regex.

### Scalability

- As the complexity of the input increases (more nested parentheses, larger molecules), all approaches show increased time and memory usage, with the regex approach scaling the worst.
