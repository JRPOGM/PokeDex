package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input	string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input: "Pokemon is amazing",
			expected: []string{"pokemon", "is", "amazing"},
		},
		{
			input : "coding is actually HARD",
			expected: []string{"coding", "is", "actually", "hard"},
		},
		{
			input: "Bungie and Sony killed Destiny on purpose",
			expected: []string{"bungie", "and", "sony", "killed", "destiny", "on", "purpose"},
		},
		{
			input: "I write POETRY like a God",
			expected: []string{"i", "write", "poetry", "like", "a", "god"},
		},
		{
			input: "One more for the Road",
			expected: []string{"one", "more", "for", "the", "road"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("input and expectation do not match: %v is not %v", actual, c.expected)
			continue
		}
		for i :=range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word is not present in both: %v is not %v", actual, c.expected)
			}
		}
	}
}