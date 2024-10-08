package cldstorage

import "github.com/cloudinary/cloudinary-go/v2"

type cloudinaryStorage struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryStorage(cld *cloudinary.Cloudinary) *cloudinaryStorage {
	return &cloudinaryStorage{cld}
}
