# Library Management System Documentation

## Overview
The Library Management System is designed to manage books and members in a library. It allows users to add, remove, borrow, and return books, as well as list available and borrowed books.

## Functions

### `getInputStr(s string) (string, error)`
Prompts the user with a message and reads a string input from the console.

- **Parameters:**
   - `s`: The prompt message to display.
- **Returns:**
   - The trimmed string input from the user.
   - An error if there is an issue reading the input.

### `getInputInt(s string) (int, error)`
Prompts the user with a message and reads an integer input from the console.

- **Parameters:**
   - `s`: The prompt message to display.
- **Returns:**
   - The integer input from the user.
   - An error if there is an issue reading the input or converting it to an integer.

### `LibraryController()`
Main controller function for the library management system. It handles user interactions and performs various operations based on user input.

- **Operations:**
   1. **Add Book**: Prompts the user to enter book details and adds the book to the library.
   2. **Remove Book**: Prompts the user to enter the book ID and removes the book from the library.
   3. **Borrow Book**: Prompts the user to enter the book ID and member ID, and borrows the book for the member.
   4. **Return Book**: Prompts the user to enter the book ID and member ID, and returns the book to the library.
   5. **List Available Books**: Lists all available books in the library.
   6. **List Borrowed Books**: Prompts the user to enter the member ID and lists all books borrowed by the member.
   7. **Exit**: Exits the library management system.

## Usage
1. Run the `LibraryController` function to start the library management system.
2. Follow the on-screen prompts to perform various operations.
3. Enter valid inputs as requested by the system to ensure smooth operation.

## Error Handling
The system will prompt the user to enter valid inputs if an error occurs during input reading or processing. Appropriate error messages are displayed for each operation if something goes wrong.

## Dependencies
- `bufio`
- `fmt`
- `os`
- `strconv`
- `strings`
- `library_management/models`
- `library_management/services`

## Models
- **Member**: Represents a library member with an ID and Name.
- **Book**: Represents a book with an ID, Title, Author, and Status.

## Services
- `NewLibrary()`: Initializes a new library instance.
- `AddBook(book Book)`: Adds a book to the library.
- `RemoveBook(id int) error`: Removes a book from the library by ID.
- `BorrowedBooks(bookID, memberID int) error`: Borrows a book for a member.
- `ReturnBook(bookID, memberID int) error`: Returns a borrowed book to the library.
- `ListAvailableBooks() []Book`: Lists all available books in the library.
- `ListBorrowedBooks(memberID int) []Book`: Lists all books borrowed by a specific member.

