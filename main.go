package main

import (
	"log"
	"os"
	"ukraine-picture/src/adapters/app"
	"ukraine-picture/src/adapters/frameworks/api"
	"ukraine-picture/src/adapters/frameworks/database"

	"github.com/joho/godotenv"
)

func initConfigs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initDb(connection string) *database.DbConnector {
	database, err := database.NewDbAdapter(connection)

	if err != nil {
		log.Fatal("Unable to connect to the database")
	}

	return database
}

func main() {
	initConfigs()
	db := initDb(os.Getenv("POSTGRES_CONNECTION"))

	mediaStorage := database.NewMediaStorage(db)
	storiesStorage := database.NewStoriesStorage(db)

	mediaService := app.NewMediaService(mediaStorage)
	storiesService := app.NewStoryService(storiesStorage)

	server := api.NewAdapter(storiesService, mediaService)

	server.Run()
}
