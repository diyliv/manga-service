package main

import (
	"log"

	"github.com/diyliv/manga-service/configs"
	"github.com/diyliv/manga-service/internal/server"
	"github.com/diyliv/manga-service/pkg/logger"
	"github.com/diyliv/manga-service/pkg/storage/mongo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error while reading .env file: %v\n", err)
	}

	// username := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	// password := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")

	cfg := configs.ReadConfig()
	mongo := mongo.InitConnect(*cfg)
	logger := logger.InitLogger(mongo)

	server := server.NewServer(logger, cfg)
	server.RungRPC()
}
