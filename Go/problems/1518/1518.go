// 1518. Water Bottles
package water

func numWaterBottles(numBottles int, numExchange int) int {
	canDrink := numBottles
	emptyBottles := numBottles
	for emptyBottles >= numExchange {
		emptyBottles -= numExchange
		canDrink++
		emptyBottles++
	}
	return canDrink
}
