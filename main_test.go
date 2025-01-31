package main

import (
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "Single word", input: "word", want: 1},
		{name: "Empty string", input: "", want: 0},
		{name: "Trailing spaces", input: "word ", want: 1},
		{name: "Two words separated by space", input: "hello world", want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := 0
			countWords(tt.input, &wc)
			if wc != tt.want {
				t.Errorf("countWords() =File: '%v', Actual %v; want %v", tt.name, wc, tt.want)
			}
		})
	}
}

func TestCountLineLength(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected int
	}{
		{"Empty string", "", 0},
		{"Single character", "a", 1},
		{"String with multiple characters", "Hello, world!", 13},
		{"String with spaces", "Hello,     world!", 17},
		{"Multiline string", "Hello,\nworld!", 13},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countLineLength(tt.line)
			if actual != tt.expected {
				t.Errorf("countLineLength() = '%s': got %v; want %v", tt.name, actual, tt.expected)
			}
		})
	}
}
