package utils

import (
	"context"
	"io"
	"log"
	"os"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio interface {
	Connect() *minio.Client
	Store(filePath string)
	Get()
}

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSsl          bool
	BucketName      string
}

func Init() *MinioConfig {
	return &MinioConfig{
		Endpoint:        os.Getenv("MINIO_ENDPOINT"),
		AccessKeyID:     os.Getenv("MINIO_ACCESS_IDKEY"),
		SecretAccessKey: os.Getenv("MINIO_SECRET_ACCESS_KEY"),
		UseSsl:          false,
		BucketName:      os.Getenv("MINIO_BUCKET_NAME"),
	}
}

func (cfg *MinioConfig) Connect() *minio.Client {
	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSsl,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return minioClient
}

func (cfg *MinioConfig) Store(objectName string) error {
	ctx := context.Background()
	m := cfg.Connect()
	// objectName := filePath
	filePath := objectName
	contentType := "image/jpeg"

	// Upload the test file with FPutObject
	_, err := m.FPutObject(ctx, cfg.BucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}

	log.Printf("Successfully uploaded %s", objectName)
	return nil
}

func (cfg *MinioConfig) Get(objectName string) {
	m := cfg.Connect()
	object, err := m.GetObject(context.Background(), cfg.BucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Fatal(err)
	}
	defer object.Close()

	localFile, err := os.Create("local-file.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer localFile.Close()

	if _, err = io.Copy(localFile, object); err != nil {
		log.Fatal(err)
	}
}
