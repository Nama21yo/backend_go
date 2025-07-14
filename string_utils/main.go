package main

import (
	"fmt"
)

func main() {
	text := "Hello, hello world! The world is big. Hello again."
	freq := WordFrequency(text)
	fmt.Println("Word Frequencies:", freq)

	palindromeText := "A man, a plan, a canal: Panama"
	fmt.Printf("Is \"%s\" a palindrome? %v\n", palindromeText, IsPalindrome(palindromeText))
}
