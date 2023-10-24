package piscine

func Itoa(n int) string {
	// Handle the special case when n is 0
	if n == 0 {
		return "0"
	}

	// Initialize variables
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	// Calculate the length of the string representation
	length := 0
	temp := n
	for temp > 0 {
		length++
		temp /= 10
	}

	// Create a byte slice to store the characters
	str := make([]byte, length)

	// Convert the number to a string in reverse order
	for i := length - 1; i >= 0; i-- {
		str[i] = byte(n%10) + '0'
		n /= 10
	}

	// Add a negative sign if necessary
	if isNegative {
		return "-" + string(str)
	}

	return string(str)
}
