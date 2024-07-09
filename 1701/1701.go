// 1701. Average Waiting Time
package restaurant

func averageWaitingTime(customers [][]int) float64 {
	var (
		totalWaitingTime int
		currentTime      int
	)

	for _, customer := range customers {
		arrivalTime := customer[0]
		timeToPrepareOrder := customer[1]

		currentTime = max(currentTime, arrivalTime)
		waitingTime := currentTime - arrivalTime + timeToPrepareOrder
		totalWaitingTime += waitingTime

		currentTime += timeToPrepareOrder
	}

	return float64(totalWaitingTime) / float64(len(customers))
}
