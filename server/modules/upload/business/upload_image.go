package business

import (
	"context"
	"mime/multipart"

	"github.com/NguyenQuy03/cinema-app/server/modules/upload/model"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type UploadImageStorage interface {
	UploadImage(ctx context.Context, c *gin.Context, file *multipart.FileHeader, param uploader.UploadParams) (*uploader.UploadResult, error)
}

type uploadImageBiz struct {
	storage UploadImageStorage
}

func NewUploadImageBiz(storage UploadImageStorage) *uploadImageBiz {
	return &uploadImageBiz{
		storage: storage,
	}
}

func (biz *uploadImageBiz) UploadMovieImage(ctx context.Context, c *gin.Context, file *multipart.FileHeader) (*uploader.UploadResult, error) {
	resp, err := biz.storage.UploadImage(c.Request.Context(), c, file, uploader.UploadParams{
		Folder: model.MovieImageFolder,
	})

	if err != nil {
		return nil, model.ErrUploadImage
	}

	return resp, nil
}
