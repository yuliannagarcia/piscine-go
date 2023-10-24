package piscine

func Join(strs []string, sep string) string {
	strNew := ""
	count := 0
	for range strs {
		count++
	}
	for i, v := range strs {
		if i == count-1 {
			strNew += v
		} else {
			strNew += v + sep
		}
	}
	return strNew
}
