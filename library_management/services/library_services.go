package services

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book, error)
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]*models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]*models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	book.Status = "Available"
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) error {
	if _, exists := l.Books[bookID]; !exists {
		return errors.New("book not found")
	}
	delete(l.Books, bookID)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status != "Available" {
		return errors.New("book is already borrowed")
	}
	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}
	book.Status = "Borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}
	found := false
	newBorrowed := []models.Book{}
	for _, b := range member.BorrowedBooks {
		if b.ID == bookID {
			found = true
		} else {
			newBorrowed = append(newBorrowed, b)
		}
	}
	if !found {
		return errors.New("book not borrowed by this member")
	}
	member.BorrowedBooks = newBorrowed
	book.Status = "Available"
	l.Books[bookID] = book
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	available := []models.Book{}
	for _, b := range l.Books {
		if b.Status == "Available" {
			available = append(available, b)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	member, exists := l.Members[memberID]
	if !exists {
		return nil, errors.New("member not found")
	}
	return member.BorrowedBooks, nil
}
