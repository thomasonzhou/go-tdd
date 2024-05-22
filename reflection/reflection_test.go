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
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Paris"},
			},
			[]string{"London", "Paris"},
		},
	}

	for _, test := range cases {

		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			assertSlicesEqual(got, test.ExpectedCalls, t)
		})
	}

	t.Run("with maps", func(t *testing.T) {

		mp := map[int]string{
			4:  "working hard",
			88: "lucky",
		}
		var got []string

		walk(mp, func(s string) {
			got = append(got, s)
		})

		assertContains("working hard", got, t)
		assertContains("lucky", got, t)
	})

	t.Run("with channels", func(t *testing.T) {
		ch := make(chan Language)
		go func() {
			ch <- Language{1, "日本語"}
			ch <- Language{3, "Français"}
			ch <- Language{2, "中文"}
			close(ch)
		}()

		var got []string
		want := []string{"日本語", "Français", "中文"}
		walk(ch, func(s string) {
			got = append(got, s)
		})

		assertSlicesEqual(got, want, t)
	})

	t.Run("with function", func(t *testing.T) {
		testFunc := func() (Language, Person) {
			return Language{5, "English"}, Person{"Feynman", Profile{69, "New York City"}}
		}

		want := []string{"English", "Feynman", "New York City"}
		var got []string

		walk(testFunc, func(s string) {
			got = append(got, s)
		})
		assertSlicesEqual(got, want, t)
	})
}

func assertSlicesEqual(got, want []string, t *testing.T) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v wanted %v", got, want)
	}
}

func assertContains(want string, got []string, t *testing.T) {
	t.Helper()
	for i := range got {
		if got[i] == want {
			return
		}
	}
	t.Errorf("expected %q to be in %v", want, got)
}
