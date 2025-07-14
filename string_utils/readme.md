# String Utilities in Go

This repository provides two utility functions written in Go:

1. Word Frequency Counter - Counts the frequency of each word in a string.
2. Palindrome Checker - Checks whether a string is a palindrome.

## Project Structure

```
string\_utils/
├── string\_utils.go       // Contains core utility functions
├── string\_utils\_test.go  // Unit tests for the functions
├── go.mod                // Go module definition
```

## Features

### Word Frequency Count

Counts how many times each word appears in a case-insensitive manner, ignoring punctuation.

**Function Signature**

```go
func WordFrequency(input string) map[string]int
```

**Example**

```go
input := "Go is fun. Go is powerful. Go Go Go!"
result := WordFrequency(input)
// Output: map[go:5 is:2 fun:1 powerful:1]
```

### Palindrome Check

Checks if a given string is a palindrome. This is done by removing spaces and punctuation, converting the string to lowercase, and checking if it reads the same backward.

**Function Signature**

```go
func IsPalindrome(input string) bool
```

**Example**

```go
IsPalindrome("A man, a plan, a canal: Panama") // true
IsPalindrome("Hello")                          // false
```

## How It Works

**Word Frequency**

- Lowercases the input
- Removes punctuation
- Splits the string into words
- Uses a map to count each word's frequency

**Palindrome Check**

- Removes non-letter characters and spaces
- Lowercases the input
- Compares the cleaned string with its reverse

## Tests

Unit tests are included for both functions.

**Run tests with:**

```bash
go test
```

### Sample Test Output

![Test Output](https://gist.github.com/user-attachments/assets/placeholder-test-output.png)

Replace the above image link with your own test result screenshot.

## Example Program

```go
package main

import (
	"fmt"
	"string_utils"
)

func main() {
	text := "Go Go Go is fun. Go is powerful!"
	fmt.Println(string_utils.WordFrequency(text))

	pal := "A man, a plan, a canal: Panama"
	fmt.Printf("Is '%s' a palindrome? %v\n", pal, string_utils.IsPalindrome(pal))
}
```

## Sample Output

![Sample Run Output](https://gist.github.com/user-attachments/assets/placeholder-run-output.png)

Replace the above image link with a screenshot of the actual output.

## Requirements

- Go 1.18 or later
