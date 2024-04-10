package profile

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {

	var input reqBody
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON input",
		})
		return
	}

	result, err := uploadImageToCloudinary(input.Base64Str, input.Name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"image_url": result.SecureURL,
	})
}

func uploadImageToCloudinary(base64Data, imageName string) (*uploader.UploadResult, error) {
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	cloudinaryService, err := cloudinary.NewFromURL(cloudinaryURL)

	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	decodedData, err := base64.StdEncoding.DecodeString(base64Data)

	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(decodedData)

	res, err := cloudinaryService.Upload.Upload(ctx, reader, uploader.UploadParams{
		PublicID: imageName,
	})

	if err != nil {
		fmt.Printf("failed to upload image, %v\n", err)
		return nil, err
	}
	return res, nil

}

type reqBody struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Base64Str string `json:"base64str"`
}
