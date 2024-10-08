package model

import (
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
)

var (
	ErrUploadImage = common.NewCustomError(errors.New("error upload image"), "unable to upload file to Cloudinary", "ERROR_UPLOAD_IMAGE")
)
