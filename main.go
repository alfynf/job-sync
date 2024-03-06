package main

import (
	"jobsync-be/configs"
	m "jobsync-be/middlewares"
	"jobsync-be/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// initialize database
	configs.InitDB()

	// inialize handler
	r := routes.New()
	m.LogMiddleware(r)

	// initialize server
	config := &configs.ServerConfig{
		Port:       os.Getenv("SERVER_PORT"),
		GinHandler: r,
	}
	s := config.Load()

	// serve
	s.ListenAndServe()
}
