package handlers

import (
	"books-api/models"
	"books-api/storage"
	"encoding/json"
	"net/http"
)

// GetBooksHandler menangani request GET /books
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books, err := storage.GetAllBooks()
	if err != nil {
		http.Error(w, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}
func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books, err := storage.GetBook(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Failed to retrieve book", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(books)
}
func DeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := storage.DeleteBook(r.PathValue("id")); err != nil {
		http.Error(w, "Failed to delete book", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// CreateBookHandler menangani request POST /books
func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := storage.AddBook(book); err != nil {
		http.Error(w, "Failed to add book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func UpdateBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Ambil ID dari parameter URL
	id := r.PathValue("id")
	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Panggil fungsi untuk update data
	err := storage.UpdateBookByID(id, updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book updated successfully"})
}
