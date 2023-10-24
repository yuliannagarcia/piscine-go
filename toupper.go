package piscine

func ToUpper(s string) string {
	var result string
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			result += string(r + ('A' - 'a'))
		} else {
			result += string(r)
		}
	}
	return result
}
