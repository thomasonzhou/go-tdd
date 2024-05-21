package reflection

import (
	"slices"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

func TestReflection(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"single string struct",
			struct {
				Work string
			}{"勉強"},
			[]string{"勉強"},
		},
		{
			"struct with non-string field",
			Profile{42, "London"},
			[]string{"London"},
		},
		{
			"nested field",
			Person{"Ben", Profile{78, "Paris"}},
			[]string{"Ben", "Paris"},
		},
	}

	for _, test := range cases {

		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !slices.Equal(got, test.ExpectedCalls) {
				t.Errorf("want %v got %v", test.ExpectedCalls, got)
			}
		})
	}

}
