// 2582. Pass the Pillow
package pillow

func PassThePillow(n int, time int) int {
	var curr = 1
	var reverse bool

	for i := 0; i < time; i++ {
		if curr == n {
			reverse = true
		} else if curr == 1 {
			reverse = false
		}

		if reverse {
			curr--
		} else {
			curr++
		}
	}

	return curr
}
