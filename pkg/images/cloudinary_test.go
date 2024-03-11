package images

import (
	"belajar-clodinary/utils/config"
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var path = "./../../.env"
var cloud Cloudinary

func init() {
	config.LoadConfig(path)
	cloudName := os.Getenv("CLOUDINARY_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cloud = NewCloudinary(cloudName, apiKey, apiSecret)
}
func TestBuildCloudinary(t *testing.T) {
	cloudName := os.Getenv("CLOUDINARY_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cloudinary := NewCloudinary(cloudName, apiKey, apiSecret)

	require.Nil(t, cloudinary.IsError)
}

func TestUploadByURL(t *testing.T) {
	imagePath := "https://uty.ac.id/img/logo.png"
	url, err := cloud.Upload(context.Background(), imagePath, "uty/wah")
	require.Nil(t, err)
	require.NotEmpty(t, url)
}

func TestUploadByBuffer(t *testing.T) {
	imagePath := "goper3.png"
	file, err := os.Open(imagePath)
	require.Nil(t, err)

	defer file.Close()

	buffer := bytes.NewBuffer(nil)

	io.Copy(buffer, file)

	url, err := cloud.Upload(context.Background(), buffer, "uty/wah")
	require.Nil(t, err)
	require.NotEmpty(t, url)
}

func TestRemove(t *testing.T) {
	path := "himatika/himatika/ristek/e3cb29f4-65c3-409c-819e-d99331d89a53"
	err := cloud.Remove(context.Background(), path)
	require.Nil(t, err)
}
