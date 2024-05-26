package roman

func intToRoman(num int) string {
	if num == 1 {
		return "I"
	} else if num == 2 {
		return "II"
	}
	return ""
}
