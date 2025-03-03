package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "ASDFASDF",
			expected: []string{"asdfasdf"},
		},
		{
			input:    "hello   world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello Mother NiCE to See you!",
			expected: []string{"hello", "mother", "nice", "to", "see", "you!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Mismatch in the lenght. Expected= %v, got=%v", len(c.expected), len(actual))

		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Mismatch in words. Expected=%s, got=%s", expectedWord, word)
			}

		}
	}

}
