package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get student name
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Get number of subjects
	var numSubjects int
	for {
		fmt.Print("Enter number of subjects: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		n, err := strconv.Atoi(input)
		if err == nil && n > 0 {
			numSubjects = n
			break
		}
		fmt.Println("Invalid input. Please enter a valid number greater than 0.")
	}

	subjects := make([]string, numSubjects)
	grades := make([]float64, numSubjects)

	for i := 0; i < numSubjects; i++ {
		// Subject name
		fmt.Printf("Enter name of subject %d: ", i+1)
		subject, _ := reader.ReadString('\n')
		subjects[i] = strings.TrimSpace(subject)

		// Grade input with validation
		for {
			fmt.Printf("Enter grade for %s (0 - 100): ", subjects[i])
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			grade, err := strconv.ParseFloat(input, 64)
			if err == nil && grade >= 0 && grade <= 100 {
				grades[i] = grade
				break
			}
			fmt.Println("Invalid grade. Must be a number between 0 and 100.")
		}
	}

	// Calculate average
	average, err := calculateAverage(grades)
	if err != nil {
		fmt.Println("Error calculating average:", err)
		return
	}

	// Output
	fmt.Println("\n📋 Grade Report")
	fmt.Printf("Student Name: %s\n", name)
	for i := 0; i < numSubjects; i++ {
		fmt.Printf("- %s: %.2f\n", subjects[i], grades[i])
	}
	fmt.Printf("🎓 Average Grade: %.2f\n", average)
}
