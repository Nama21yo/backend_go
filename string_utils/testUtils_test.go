package main

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	input := "Go is fun. Go is powerful. Go Go Go!"
	expected := map[string]int{
		"go":       5,
		"is":       2,
		"fun":      1,
		"powerful": 1,
	}
	result := WordFrequency(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"racecar", true},
		{"Was it a car or a cat I saw?", true},
		{"Not a palindrome", false},
		{"", true},
	}

	for _, tt := range tests {
		result := IsPalindrome(tt.input)
		if result != tt.expected {
			t.Errorf("IsPalindrome(%q) = %v; expected %v", tt.input, result, tt.expected)
		}
	}
}
