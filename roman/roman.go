package roman

func intToRoman(num int) (roman string) {
	if num <= 3 {
		for i := 0; i < num; i++ {
			roman += "I"
		}
	}
	return roman
}
