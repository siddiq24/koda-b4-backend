package libs

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

func SaveUploadedFile(c *gin.Context, file multipart.File, header *multipart.FileHeader, nama string) (string, error) {
	defer file.Close()
	var filename string

	fileExt := filepath.Ext(header.Filename)
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
		return "", fmt.Errorf("format image is wrong")
	}
	if header.Size > (1024 * 1024) {
		return "", fmt.Errorf("Ukuran image terlalu besar")
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	filename = fmt.Sprintf("%s%s", nama, fileExt)
	os.MkdirAll("images/user", os.ModePerm)

	localPath := fmt.Sprintf("images/user/%s", filename)

	resized := imaging.Resize(img, 1000, 0, imaging.Lanczos)
	if err := imaging.Save(resized, localPath); err != nil {
		fmt.Println(err)
		return "", err
	}
	return filename, nil
}
