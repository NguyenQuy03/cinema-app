package cldstorage

import (
	"context"
	"mime/multipart"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func (storage *cloudinaryStorage) UploadImage(ctx context.Context, c *gin.Context, file *multipart.FileHeader, param uploader.UploadParams) (*uploader.UploadResult, error) {
	resp, err := storage.cld.Upload.Upload(c.Request.Context(), file, param)

	if err != nil {
		return nil, common.ErrDB(err)
	}

	return resp, nil
}
