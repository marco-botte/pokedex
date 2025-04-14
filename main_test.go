package main

import (
	"reflect"
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
			input:    "go is    fun ",
			expected: []string{"go", "is", "fun"},
		},
		{
			input:    "   catch PIKaCHU",
			expected: []string{"catch", "pikachu"},
		},
		{
			input:    "     ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("cleanInput(%q) = %v; want %v", c.input, actual, c.expected)
		}
	}
}
