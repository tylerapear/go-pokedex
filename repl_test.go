package main

import "testing"

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
			input: "helloworld ",
			expected: []string{"helloworld"},
		},
		{
			input: " hello world again ",
			expected: []string{"hello", "world", "again"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("TestCleanInput failed: for input '%s' expected %d words but got %d", c.input, len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("TestCleanInput failed: for input '%s' expected word %d to be '%v' but got '%v'", c.input, i, expectedWord, word)
			}
		}
	}
}