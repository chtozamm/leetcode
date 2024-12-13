// 1460. Make Two Arrays Equal by Reversing Subarrays
package equal_arrays

func canBeEqual(target []int, arr []int) bool {
	hashTable := make(map[int]int, len(arr))

	for _, n := range arr {
		hashTable[n]++
	}

	for _, n := range target {
		if hashTable[n] == 0 {
			return false
		}

		hashTable[n]--
	}

	return true
}
