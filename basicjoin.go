package piscine

func BasicJoin(strs []string) string {
	strNew := ""
	for _, v := range strs {
		strNew += v
	}
	return strNew
}
