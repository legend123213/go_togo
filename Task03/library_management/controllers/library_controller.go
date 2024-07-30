package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

func getInputStr(s string) (string, error) {
	 fmt.Print(s)
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    return strings.TrimSpace(input),err
}
func getInputInt(s string) (int,error){
	 fmt.Print(s)
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
	 num,err_ := strconv.Atoi(strings.TrimSpace(input))
    return num,err_
}

func LibraryController(){
 lib := services.NewLibrary()
 memberID,err_id := getInputInt("Enter your id: ")
 nameMember,err_title := getInputStr("Enter Your Name: ")
 if err_id != nil || err_title != nil {
	panic("Please Enter Vaild Input")
 }
 var member models.Member
 member.ID = memberID
 member.Name = nameMember
 lib.Members[memberID] = member

 for{
		fmt.Println("\n=== Library Management System ===")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		choice, _ := getInputInt("Please Enter your choice: ")
		switch choice {

		case 1:
			   id,err_id := getInputInt("Enter Book ID: ")
            title,err_title := getInputStr("Enter Book Title: ")
            author,err_author := getInputStr("Enter Book Author: ")
				if (err_id!=nil || err_title!=nil || err_author!=nil){
					fmt.Println("🚨Somethings wrong adding book not successful")
					continue
				}
            book := models.Book{ID: id, Title: title, Author: author, Status: "Available"}
            lib.AddBook(book)
				fmt.Println("🟢Book Add successful")
		case 2:
			id,err_ := getInputInt("Enter Book ID to remove: ")
			if err_ !=nil{
				fmt.Println("🚨Somethings wrong adding book not successful")
				continue
			}
         err := lib.RemoveBook(id)
			if err == nil{
				fmt.Println("🟢Book removed successfully!")
			}else{
				fmt.Println("🚨",err)
			}
         
		case 3:
			bookID,err_book:= getInputInt("Enter Book ID to borrow: ")
            memberID,err_member := getInputInt("Enter Member ID: ")
				if err_book != nil || err_member != nil {
					fmt.Println("🚨Somethings wrong adding book not successful")
					continue
				}
            err := lib.BorrowedBooks(bookID, memberID)
            if err != nil {
                fmt.Println("🚨Error borrowing book:", err)
            } else {
                fmt.Println("🟢Book borrowed successfully!")
            }
        case 4:
            bookID,err_book := getInputInt("Enter Book ID to return: ")
            memberID,err_member := getInputInt("Enter Member ID: ")
				if err_book != nil || err_member != nil {
					fmt.Println("🚨Somethings wrong adding book not successful")
					continue
				}
            err := lib.ReturnBook(bookID, memberID)
            if err != nil {
                fmt.Println("🚨Error returning book:", err)
            } else {
                fmt.Println("Book returned successfully!")
            }
        case 5:
            availableBooks := lib.ListAvailableBooks()
            fmt.Println("\nAvailable Books:")
            for _, book := range availableBooks {
                fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
            }
        case 6:
            memberID,err := getInputInt("Enter Member ID: ")
            borrowedBooks := lib.ListBorrowedBooks(memberID)
				if err != nil {
					fmt.Println("🚨Somethings wrong adding book not successful")
					continue
				}
            fmt.Println("\nBorrowed Books:")
            for _, book := range borrowedBooks {
                fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
            }
        case 7:
            fmt.Println("👋Exiting the system. Goodbye!")
            return
        default:
				fmt.Println("🚨Please Enter Valid input")
		}

 } 

}