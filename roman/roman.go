package roman

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(arabicNum int) string {

	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabicNum >= numeral.Value {
			arabicNum -= numeral.Value
			result.WriteString(numeral.Symbol)
		}
	}
	return result.String()
}
