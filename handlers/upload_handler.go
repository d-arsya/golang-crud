package handlers

import (
	"books-api/storage"
	"fmt"
	"net/http"
)

// UploadFileHandler menangani upload file ke Supabase
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // Max 10MB file size

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath, err := storage.UploadFileToSupabase(file, handler.Filename)
	if err != nil {
		http.Error(w, "Upload failed", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", filePath)
}
