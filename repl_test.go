package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"Charmander", "Bulbasaur", "PIKACHU"},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf(`--------------------------------- 
						Incorrect Lenght:
						Inputs:     (%v)
						Expecting:  %v
						Actual:     %v
						Fail
						`, len(c.input), len(c.expected), len(actual))
		} else {
			fmt.Printf(`---------------------------------
				Inputs:     (%v)
				Expecting:  %v
				Actual:     %v
				Pass
				`, c.input, c.expected, actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf(`---------------------------------
						Words do not match:
						
						Expecting:  %v
						Actual:     %v
						Fail
						`, expectedWord, word)
			} else {
				fmt.Printf(`---------------------------------
				Expecting:  %v
				Actual:     %v
				Pass
				`, expectedWord, word)
			}
		}

	}
}
