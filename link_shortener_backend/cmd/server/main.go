package main

import (
	"log"
	"log/slog"
	"os"

	"linkShortener/configs"
	"linkShortener/internal/database"
	"linkShortener/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	configData := configs.LoadEnvVariables()
	db, dbCloser, err := database.ConnectToDB(*configData)
	defer dbCloser()

	if err != nil {
		log.Fatal("Db connection error: ", err)
	}

	counter, counterCloser, err := database.ConnectCounter(*configData)
	defer counterCloser()

	if err != nil {
		log.Fatal("Counter connection error: ", err)
	}

	err = database.Migrate(db)

	if err != nil {
		log.Fatal("Db migration error: ", err)
	}

	//createLinkUseCase := aservice.NewCreateLinkUseCase(db, logger)
	//controller := handler.NewLinkController(createLinkUseCase, logger)

	router := gin.Default()

	router.POST("/api/v1/links", handler.CreateLink)
	router.GET("/api/v1/links/:code", handler.Forward)
	router.GET("/api/v1/links/:code/counter", handler.GetCount)

	//todo: graceful shutdown?
	err = router.Run()
	if err != nil {
		log.Fatalf("Server run error: %v", err)
	}
}
