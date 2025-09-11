package main

import (
	"linkShortener/configs"
	"linkShortener/internal/database"
	"linkShortener/internal/handler"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.Load("./deployment/.env")
	database.ConnectToDB()
	database.Migrate()
}

func main() {
	router := gin.Default()

	router.POST("/api/v1/links", handler.CreateLink)
	router.GET("/api/v1/links/:code", handler.Forward)

	router.Run()
}
