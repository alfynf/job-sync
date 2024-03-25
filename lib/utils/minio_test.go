package utils

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var CFG *MinioConfig

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	CFG = &MinioConfig{
		Endpoint:        os.Getenv("MINIO_ENDPOINT"),
		AccessKeyID:     os.Getenv("MINIO_ACCESS_IDKEY"),
		SecretAccessKey: os.Getenv("MINIO_SECRET_ACCESS_KEY"),
		UseSsl:          false,
		BucketName:      os.Getenv("MINIO_BUCKET_NAME"),
	}
	m.Run() // eksekusi semua unit test
}

func TestInitMinio(t *testing.T) {
	CFG.Connect()
}

func TestStoreMinio(t *testing.T) {
	CFG.Store("nami.jpg")
}

func TestGetMinio(t *testing.T) {
	CFG.Get("nami.jpg")
}
