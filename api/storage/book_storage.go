package storage

import (
	"books-api/api/config"
	"books-api/api/models"
	"context"
	"errors"
	"strconv"
)

// GetAllBooks mengambil semua data buku dari database
func GetAllBooks() ([]models.Book, error) {
	rows, err := config.DB.Query(context.Background(), "SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
func GetBook(id string) (*models.Book, error) {
	bookId, _ := strconv.Atoi(id)
	var book models.Book
	err := config.DB.QueryRow(context.Background(), "SELECT * FROM books WHERE id = $1", bookId).
		Scan(&book.ID, &book.Title, &book.Author)

	if err != nil {
		return nil, errors.New("book not found")
	}

	return &book, nil
}

func UpdateBookByID(id string, updatedBook models.Book) error {
	bookId, _ := strconv.Atoi(id)
	result, err := config.DB.Exec(
		context.Background(),
		"UPDATE books SET title = $1, author = $2 WHERE id = $3",
		updatedBook.Title, updatedBook.Author, bookId,
	)

	if err != nil {
		return err
	}

	// Cek apakah ada baris yang terpengaruh
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("book not found")
	}

	return nil
}

// AddBook menambahkan buku ke database
func AddBook(book models.Book) error {
	_, err := config.DB.Exec(context.Background(),
		"INSERT INTO books (title, author) VALUES ($1, $2)", book.Title, book.Author)
	return err
}
func DeleteBook(id string) error {
	bookId, _ := strconv.Atoi(id)
	_, err := config.DB.Exec(context.Background(),
		"DELETE FROM books WHERE id=$1", bookId)
	return err
}
