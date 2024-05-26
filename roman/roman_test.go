package roman

import (
	"fmt"
	"testing"
)

func TestRomanConverter(t *testing.T) {

	cases := []struct {
		ArabicNum   int
		RomanSymbol string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{14, "XIV"},
		{18, "XVIII"},
		{20, "XX"},
		{39, "XXXIX"},
		{40, "XL"},
		{47, "XLVII"},
		{49, "XLIX"},
		{50, "L"},
		{90, "XC"},
		{100, "C"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{3000, "MMM"},
	}

	for _, test := range cases {
		description := fmt.Sprintf("%v to %q", test.ArabicNum, test.RomanSymbol)
		t.Run(description, func(t *testing.T) {
			got := intToRoman(test.ArabicNum)

			assertStringEqual(got, test.RomanSymbol, t)
		})
	}
}

func assertStringEqual(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}
