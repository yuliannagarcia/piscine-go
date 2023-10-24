package piscine

func RecursivePower(nb, power int) int {
	if power < 0 {
		return 0 // Negative powers return 0
	}

	if power == 0 {
		return 1 // Any number raised to the power of 0 is 1
	}

	return nb * RecursivePower(nb, power-1)
}
