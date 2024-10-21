package db

import (
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func InitCloudinaryStorage() (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse Redis URL: %w", err)
	}

	return cld, nil
}
