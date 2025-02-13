package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// UploadFileToSupabase mengunggah file ke Supabase Storage
func UploadFileToSupabase(file multipart.File, fileName string) (string, error) {
	url := fmt.Sprintf("%s/storage/v1/object/%s/%s", os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_BUCKET"), fileName)
	req, err := http.NewRequest("PUT", url, file)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("SUPABASE_KEY"))
	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Failed to upload file: %s", string(body))
	}

	responseMap := map[string]string{}
	json.NewDecoder(resp.Body).Decode(&responseMap)

	return responseMap["Key"], nil
}
