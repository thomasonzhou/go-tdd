package roman

import "testing"

func TestRomanConverter(t *testing.T) {

	cases := []struct {
		description string
		num         int
		want        string
	}{
		{"1 to I",
			1,
			"I",
		},
		{"2 to II",
			2,
			"II",
		},
		{"3 to III",
			3, "III"},
		{"4 to IV", 4, "IV"},
		{"5 to V", 5, "V"},
	}

	for _, test := range cases {
		t.Run(test.description, func(t *testing.T) {
			got := intToRoman(test.num)

			assertStringEqual(got, test.want, t)
		})
	}
}

func assertStringEqual(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}
