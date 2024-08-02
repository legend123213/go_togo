package services

import (
	"errors"
	"library_management/models"
)

type Library_Manager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowedBooks(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

// AddBook adds a book to the library
func (lib *Library) AddBook(book models.Book) {
	lib.Books[book.ID] = book
}

// RemoveBook removes a book from the library
func (lib *Library) RemoveBook(id int) error{
	_,ok :=lib.Books[id]
	if !ok{
		return errors.New("No Book Found")
	}
	delete(lib.Books, id)
	return nil
}

// BorrowedBooks borrows a book from the library
func (lib *Library) BorrowedBooks(bookID int, memberID int) error {
	book, err := lib.Books[bookID]
	member := lib.Members[memberID]
	if !err {
		return errors.New("book is not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is unavailable")
	}
	
	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	lib.Books[bookID] = book
	lib.Members[memberID] = member
	return nil
}

// ReturnBook returns a borrowed book to the library
func (lib *Library) ReturnBook(bookID int, memberID int) error {
	book, err := lib.Books[bookID]
	member, exist := lib.Members[memberID]
	if !err {
		return errors.New("the book is not from this library")
	}
	if !exist {
		return errors.New("member not found")
	}
	for i, book := range member.BorrowedBooks {
		if bookID == book.ID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}
	book.Status = "Available"
	lib.Books[bookID] = book
	lib.Members[memberID] = member
	return nil
}

// ListAvailableBooks returns a list of available books in the library
func (lib *Library) ListAvailableBooks() []models.Book {
	listAvailableBook := []models.Book{}
	for _, book := range lib.Books {
		if book.Status == "Available" {
			listAvailableBook = append(listAvailableBook, book)
		}
	}
	return listAvailableBook
}

// ListBorrowedBooks returns a list of books borrowed by a member
func (lib *Library) ListBorrowedBooks(memberID int) []models.Book {
	borrowed, exist := lib.Members[memberID]
	if !exist {
		return nil
	}
	return borrowed.BorrowedBooks
}
