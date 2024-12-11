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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  TeSSt Now forever  ",
			expected: []string{"tesst", "now", "forever"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("Length of input and expected length do not match.")
            return
        }
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("input[%d] -> '%s' and expected[%d] -> '%s' do not match.", i, word, i, expectedWord)
                return
            }
		}
	}
}
