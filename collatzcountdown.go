package piscine

func CollatzCountdown(start int) int {
	step := 0

	if start <= 0 {
		return -1 // Invalid input
	}

	for start != 1 {
		if start%2 == 0 {
			start /= 2
		} else {
			start = 3*start + 1
		}
		step++
	}

	return step
}
