package main

import (
	"database/sql"
	"log"

	"linkShortener/configs"
	"linkShortener/internal/database"
	"linkShortener/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	//todo: init logger
	logger := log.DefaultJson() // GRPC | SENTRY
	db := database.MONGO()
	createLinkUseCase := aservice.NewCreateLinkUseCase(db, logger)
	controller := handler.NewLinkController(createLinkUseCase, logger)

	//todo: init root context
	//todo: return err,  from internal packages and handle it here
	config, err := configs.Load() // for local development
	if err != nil {
		logger.Fatal("Config load error: ", err)
	}
	//todo: why redis and postgres together?
	//todo: pass context to db connect, return error, return closer
	db, closer, err := database.ConnectToDB(config)
	if err != nil {
		logger.Fatal("Db connection error: ", err)
	}
	defer closer()
	//todo: pass context to db connect, return error, return closer
	counter, closer, err := database.ConnectCounter(config)
	if err != nil {
		logger.Fatal("Db connection error: ", err)
	}
	defer closer()

	// todo: error from migrate?
	database.Migrate(db)

	router := gin.Default()

	router.POST("/api/v1/links", handler.CreateLink)
	router.GET("/api/v1/links/:code", handler.Forward)
	router.GET("/api/v1/links/:code/counter", handler.GetCount)

	//todo: graceful shutdown?
	router.Run()

	//todo: error processing, return closer from connect?
	postgresDb, _ := db.DB()
	//todo: to separate defers?
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
