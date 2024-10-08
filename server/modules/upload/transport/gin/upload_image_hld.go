package ginTrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/modules/upload/business"
	cldstorage "github.com/NguyenQuy03/cinema-app/server/modules/upload/storage/cloudinary"

	"github.com/cloudinary/cloudinary-go/v2"

	"github.com/gin-gonic/gin"
)

func UploadMovieImage(cld *cloudinary.Cloudinary) func(*gin.Context) {
	return func(c *gin.Context) {

		// Get the file from the request
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
			return
		}

		cloudinary := cldstorage.NewCloudinaryStorage(cld)
		business := business.NewUploadImageBiz(cloudinary)

		resp, err := business.UploadMovieImage(c.Request.Context(), c, file)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)

			return
		}

		// Respond with the URL of the uploaded image
		c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "url": resp.SecureURL})
	}
}
