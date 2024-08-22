// 134. Gas Station
package gas_station

func canCompleteCircuit(gas []int, cost []int) int {
	var fuel, globalFuel, start int

	for i := range gas {
		globalFuel += gas[i] - cost[i]
		fuel += gas[i] - cost[i]

		if fuel < 0 {
			start = i + 1
			fuel = 0
		}
	}

	if globalFuel < 0 {
		return -1
	}

	return start
}
