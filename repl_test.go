package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{{
		input:    "    hello world  ",
		expected: []string{"hello", "world"},
	}, {
		input:    "   Hello  World",
		expected: []string{"hello", "world"},
	}, {
		input:    "   Charmander Bulbasaur PIKACHU  ",
		expected: []string{"charmander", "bulbasaur", "pikachu"},
	}}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length not match %v vs %v", actual, c.expected)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Actual word: %v and expected word:%v do not match.", actual[i], c.expected[i])
			}
		}
	}
}
