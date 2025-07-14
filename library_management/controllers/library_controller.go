package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	// "strings"
)

func StartConsole(library *services.Library) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nLibrary Menu:")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter Book ID: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Enter Title: ")
			scanner.Scan()
			title := scanner.Text()
			fmt.Print("Enter Author: ")
			scanner.Scan()
			author := scanner.Text()
			library.AddBook(models.Book{ID: id, Title: title, Author: author})
			fmt.Println("Book added.")
		case "2":
			fmt.Print("Enter Book ID to remove: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			err := library.RemoveBook(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book removed.")
			}
		case "3":
			fmt.Print("Enter Book ID to borrow: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Enter Member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())
			if _, exists := library.Members[memberID]; !exists {
				fmt.Print("Enter Member Name: ")
				scanner.Scan()
				name := scanner.Text()
				library.Members[memberID] = &models.Member{ID: memberID, Name: name}
			}
			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed.")
			}
		case "4":
			fmt.Print("Enter Book ID to return: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Enter Member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())
			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned.")
			}
		case "5":
			fmt.Println("Available Books:")
			for _, b := range library.ListAvailableBooks() {
				fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
			}
		case "6":
			fmt.Print("Enter Member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())
			books, err := library.ListBorrowedBooks(memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Books borrowed by member %d:\n", memberID)
				for _, b := range books {
					fmt.Printf("ID: %d | Title: %s\n", b.ID, b.Title)
				}
			}
		case "7":
			fmt.Println("Exiting.")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}
