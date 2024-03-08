package configs

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type configDB struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         string
	SslMode      string
}

var DB *gorm.DB

func InitDB() *gorm.DB {
	c := &configDB{
		Host:         os.Getenv("DB_HOST"),
		Username:     os.Getenv("DB_USERNAME"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
		Port:         os.Getenv("DB_PORT"),
		SslMode:      os.Getenv("DB_SSL_MODE"),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Bangkok",
		c.Host,
		c.Username,
		c.Password,
		c.DatabaseName,
		c.Port,
		c.SslMode,
	)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return DB

}
