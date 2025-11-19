package libs

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveUploadedFile(c *gin.Context, header *multipart.FileHeader, folder string) (string, error) {
	ext := filepath.Ext(header.Filename)
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
		return "", fmt.Errorf("format image salah")
	}

	if header.Size > (5 * 1024 * 1024) {
		return "", fmt.Errorf("file terlalu besar (max 5MB)")
	}

	// Pastikan folder ada
	os.MkdirAll(folder, os.ModePerm)

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	path := filepath.Join(folder, filename)

	// Simpan file
	if err := c.SaveUploadedFile(header, path); err != nil {
		return "", err
	}

	return path, nil
}
