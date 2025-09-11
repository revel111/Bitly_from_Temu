package main

import (
	"linkShortener/configs"
	"linkShortener/internal/database"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.Load("./deployments/.env")
	database.ConnectToDB()
	database.Migrate()
}

func main() {
	server := gin.Default()

	server.Run()
}
