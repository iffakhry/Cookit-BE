package storage

import (
	"context"
	"io"
	"log"
	"mime/multipart"
	"strings"

	"alta-cookit-be/app/config"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var (
	DEFAULT_GCS_LINK string = "https://storage.googleapis.com/altabucket/"
)

type ClientUploader struct {
	storageClient *storage.Client
	projectID     string
	bucketName    string
	uploadPath    string
}

var clientUploader *ClientUploader

func GetStorageClient() *ClientUploader {
	if clientUploader == nil {
		client, err := storage.NewClient(context.Background(), option.WithoutAuthentication())
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}

		clientUploader = &ClientUploader{
			storageClient: client,
			bucketName:    config.GCP_BUCKET_NAME,
			projectID:     config.GCP_PROJECT_ID,
			uploadPath:    "static/images/",
		}

		return clientUploader
	}
	return clientUploader
}

// UploadFile uploads an object
func (c *ClientUploader) UploadFile(file multipart.File, objectName string) (fileLocation string, err error) {
	ctx := context.Background()

	// Upload an object with storage.Writer.
	wc := c.storageClient.Bucket(c.bucketName).Object(c.uploadPath + objectName).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	return DEFAULT_GCS_LINK + wc.Name, nil
}

func (c *ClientUploader) DeleteFile(objectName string) error {
	ctx := context.Background()

	wc := c.storageClient.Bucket(c.bucketName).Object(strings.Replace(objectName, DEFAULT_GCS_LINK, "", 1))
	if err := wc.Delete(ctx); err != nil {
		return err
	}

	return nil
}
