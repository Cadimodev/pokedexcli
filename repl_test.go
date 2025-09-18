package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/Cadimodev/pokedexcli/internal/pokecache"
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
			expected: []string{"charmander", "bulbasaur", "pikachu"},
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

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	exampleURL := "https://example.com"
	cache.Add(exampleURL, []byte("testdata"))

	_, ok := cache.Get(exampleURL)
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(exampleURL)
	if ok {
		t.Errorf("expected not to find key")
	}
}
