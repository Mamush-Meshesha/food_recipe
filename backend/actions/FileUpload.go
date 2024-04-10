package actions

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

func FileUpload(c *gin.Context) {

	var input reqBody
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON input",
		})
		return
	}
	var imageURLs []string
	for _, image := range input.Images {
		result, err := uploadImageToCloudinary(image.Base64Str, image.Name)
		if err != nil {
			fmt.Printf("Error uploading image %s: %v\n", image.Name, err)
			continue
		}
		imageURLs = append(imageURLs, result.SecureURL)
	}
	c.JSON(http.StatusOK, gin.H{
		"image_urls": imageURLs,
	})
}

func uploadImageToCloudinary(base64Data, imageName string) (*uploader.UploadResult, error) {
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	cloudinaryService, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	// Decode the base64 data
	decodedData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(decodedData)

	// Upload the image to Cloudinary
	res, err := cloudinaryService.Upload.Upload(ctx, reader, uploader.UploadParams{
		PublicID: imageName,
	})

	if err != nil {
		fmt.Printf("Failed to upload file, %v\n", err)
		return nil, err
	}

	return res, nil
}

type reqBody struct {
	Images []struct {
		Name      string `json:"name"`
		Base64Str string `json:"base64str"`
		Type      string `json:"type"`
	} `json:"images"`
}
