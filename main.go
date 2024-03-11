package main

import (
	"belajar-clodinary/pkg/images"
	"belajar-clodinary/utils/config"
	"context"
	"os"
)

type Cloudsvc interface {
	Upload(ctx context.Context, file interface{}, path string) (string, error)
	Remove(ctx context.Context, path string) error
}

type Services struct {
	cloud Cloudsvc
}

func main() {
	config.LoadConfig(".env")
	cloudName := os.Getenv("CLOUDINARY_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	image := images.NewImage(cloudName, apiKey, apiSecret)

	svc := Services{
		cloud: image.BuildCloudinary(),
	}

	svc.cloud.Upload(context.Background(), "", "")
}
