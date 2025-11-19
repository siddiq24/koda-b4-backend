package libs

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(localPath string) (string, error) {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return "", fmt.Errorf("cloudinary init fail: %v", err)
	}

	resp, err := cld.Upload.Upload(context.Background(), localPath, uploader.UploadParams{
		PublicID: fmt.Sprintf("products/%d", time.Now().UnixNano()),
	})

	// Hapus file lokal setelah upload
	os.Remove(localPath)

	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}
