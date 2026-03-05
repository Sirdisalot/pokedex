package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charamander Bulbasaur PIKACHU",
			expected: []string{"charamander", "bulbasaur", "pikachu"},
		},
		{
			input: "   singleWord  ",
			expected: []string{"singleword"},
		},
		{
			input: "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, word, expectedWord)
				continue
			}
		}
	}

}