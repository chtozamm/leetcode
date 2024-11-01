// 2418. Sort the People
package people_heights

import "slices"

func sortPeople(names []string, heights []int) []string {
	// Map heights to names
	heightMap := make(map[int]string)
	for i, height := range heights {
		heightMap[height] = names[i]
	}

	// Make a copy of heights to prevent modifying the original slice
	sortedHeights := make([]int, len(heights))
	copy(sortedHeights, heights)

	// Sort heights in descending order
	slices.Sort(sortedHeights)
	slices.Reverse(sortedHeights)

	res := make([]string, 0, len(names))

	for _, height := range sortedHeights {
		res = append(res, heightMap[height])
	}

	return res
}
