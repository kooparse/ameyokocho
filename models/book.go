// All methods to write/read in the database related to
// the Book model are here.
package models

import (
	"github.com/kooparse/ameyokocho/utils"
	"log"
)

type Book struct {
	Title         string
	Description   string
	ISBN_10       string
	ISBN_13       string
	Language      string
	PublishedDate string
	Authors       []string
}

func AddBook(book Book) {
	sqlQuery := `
    INSERT INTO 
      books(
        title, 
        description, 
        isbn_10, 
        isbn_13, 
        language, 
        published_date
      ) 
    VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := DB.Exec(
		sqlQuery,
		book.Title,
		book.Description,
		book.ISBN_10,
		book.ISBN_13,
		book.Language,
		book.PublishedDate,
	)

	utils.ErrorCheck(err)
	log.Printf("Book '%s' added to database\n", book.Title)
}
