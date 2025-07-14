package main

import (
	"strings"
	"unicode"
)

// WordFrequency returns a map with the frequency of each word (case-insensitive, no punctuation).
func WordFrequency(text string) map[string]int {
	wordFreq := make(map[string]int)
	var wordBuilder strings.Builder

	for _, char := range text {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			wordBuilder.WriteRune(unicode.ToLower(char))
		} else if wordBuilder.Len() > 0 {
			word := wordBuilder.String()
			wordFreq[word]++
			wordBuilder.Reset()
		}
	}
	if wordBuilder.Len() > 0 {
		word := wordBuilder.String()
		wordFreq[word]++
	}
	return wordFreq
}

// IsPalindrome returns true if the input string is a palindrome (ignores case, spaces, punctuation).
func IsPalindrome(s string) bool {
	var filtered []rune
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			filtered = append(filtered, unicode.ToLower(char))
		}
	}

	for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
		if filtered[i] != filtered[j] {
			return false
		}
	}
	return true
}
