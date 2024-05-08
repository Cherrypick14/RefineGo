package main

import (
	"reloaded/functions"
	"strings"
	"testing"
)

func TestCapitalize(t *testing.T) {
	cap := []struct {
		input    string
		expected string 
	}{
		{"it (cap)", "It"},
		{"it was the age of foolishness (cap, 6)","It Was The Age Of Foolishness"},
	}

	for _, char := range cap {
		output := functions.Capitalize(strings.Split(char.input, " "))
		if output != char.expected {
			t.Errorf("Expected: %s Output: %s\n", char.expected, output)
		}
	}
}
