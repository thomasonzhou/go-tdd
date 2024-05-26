package roman

import "strings"

func intToRoman(arabicNum int) string {

	var result strings.Builder
	for arabicNum > 0 {
		switch {
		case arabicNum > 9:
			arabicNum -= 10
			result.WriteString("X")
		case arabicNum > 8:
			arabicNum -= 9
			result.WriteString("IX")
		case arabicNum > 4:
			arabicNum -= 5
			result.WriteString("V")
		case arabicNum > 3:
			arabicNum -= 4
			result.WriteString("IV")
		default:
			arabicNum -= 1
			result.WriteString("I")
		}
	}
	return result.String()
}
