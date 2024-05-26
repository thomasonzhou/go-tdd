package roman

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
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
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestArabicToRoman(t *testing.T) {
	for _, test := range cases {
		description := fmt.Sprintf("%v to %q", test.Arabic, test.Roman)
		t.Run(description, func(t *testing.T) {
			got := ArabicToRoman(test.Arabic)
			assertStringEqual(got, test.Roman, t)
		})
	}
}
func TestRomanToArabic(t *testing.T) {
	for _, test := range cases[:4] {
		description := fmt.Sprintf("%v to %q", test.Arabic, test.Roman)
		t.Run(description, func(t *testing.T) {
			got := RomanToArabic(test.Roman)
			assertUInt16Equal(got, test.Arabic, t)
		})
	}
}

func TestPropertiesofConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			log.Println(arabic)
			return true
		}
		roman := ArabicToRoman(arabic)
		fromRoman := RomanToArabic(roman)
		return arabic == fromRoman
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}

func assertStringEqual(got, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}

func assertUInt16Equal(got, want uint16, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}
