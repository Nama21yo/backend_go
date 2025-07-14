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


