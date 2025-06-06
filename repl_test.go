package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input		string
		expected	[]string
	}{
		{
			input:		"  hello  world  ",
			expected:	[]string{"hello", "world"},
		},
		{
			input:		"Charmander Bulbasaur PIKACHU",
			expected:	[]string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:		" I aM a TeSt StRiNg ",
			expected:	[]string{"i", "am", "a", "test", "string"},
		},
		{
			input:		"",
			expected:	[]string{},
		},
		{
			input:		"      ",
			expected:	[]string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths don't match: %d(actual) != %d(expected)", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Mismatched word: %s(actual) != %s(expected)", word, expectedWord)
			}
		}
	}
}