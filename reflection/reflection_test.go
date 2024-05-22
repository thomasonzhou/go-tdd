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

type Language struct {
	Ability int
	Name    string
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
		{
			"pointers to structs",
			&Profile{1000, "Japan"},
			[]string{"Japan"},
		},
		{
			"pure string variable",
			"what",
			[]string{"what"},
		},
		{
			"non string slices",
			[]int{2, 3, 5, 7},
			[]string{},
		},
		{
			"struct slices",
			[]Language{
				{1, "日本語"},
				{3, "Français"},
				{2, "中文"},
			},
			[]string{
				"日本語", "Français", "中文"},
		},
		{
			"pure non-string variable",
			false,
			[]string{},
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
