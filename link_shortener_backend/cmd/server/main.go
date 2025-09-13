package main

import (
	"database/sql"
	"linkShortener/configs"
	"linkShortener/internal/database"
	"linkShortener/internal/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.Load("./deployment/.env") // for local development
	db := database.ConnectToDB()
	counter := database.ConnectCounter()
	database.Migrate(db)

	router := gin.Default()

	router.POST("/api/v1/links", handler.CreateLink)
	router.GET("/api/v1/links/:code", handler.Forward)
	router.GET("/api/v1/links/:code/counter", handler.GetCount)

	router.Run()

	postgresDb, _ := db.DB()
	defer func(postgresDb *sql.DB) {
		err := postgresDb.Close()
		if err != nil {
			log.Println("Db connection close error: ", err)
			return
		}

		err = counter.Close()
		if err != nil {
			log.Println("Db connection close error: ", err)
			return
		}
	}(postgresDb)
}
