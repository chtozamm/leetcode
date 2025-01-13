package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	planted := 0
	length := len(flowerbed)

	for i := 0; i < length; i++ {
		if flowerbed[i] == 0 {
			leftPlotEmpty := i == 0 || flowerbed[i-1] == 0
			rightPlotEmpty := i == length-1 || flowerbed[i+1] == 0

			if leftPlotEmpty && rightPlotEmpty {
				flowerbed[i] = 1
				planted++
				i++ // Skip the next plot
			}
		}

		if planted >= n {
			return true
		}
	}

	return planted >= n
}
