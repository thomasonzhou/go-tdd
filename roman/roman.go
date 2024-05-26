package roman

import "strings"

func intToRoman(arabicNum int) string {

	var result strings.Builder
	for i := arabicNum; i > 0; i-- {
		if arabicNum == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}
	return result.String()
}
