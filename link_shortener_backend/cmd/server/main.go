package main

import (
	"context"
	"linkShortener/configs"
	"linkShortener/internal/handler/web"
	counterconnection "linkShortener/internal/infrastructure/counter/connection"
	"linkShortener/internal/infrastructure/counter/redis"
	linkconnection "linkShortener/internal/infrastructure/link/connection"
	"linkShortener/internal/infrastructure/link/postgres"
	usecasecounter "linkShortener/internal/usecase/counter"
	usecaselink "linkShortener/internal/usecase/link"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	configData := configs.LoadEnvVariables()
	db, dbCloser, err := linkconnection.ConnectToDB(*configData)
	defer dbCloser()

	if err != nil {
		log.Fatal("Db connection error: ", err)
	}

	counter, counterCloser, err := counterconnection.ConnectCounter(*configData)
	defer counterCloser()

	if err != nil {
		log.Fatal("Counter connection error: ", err)
	}

	err = linkconnection.Migrate(db)
	if err != nil {
		log.Fatal("Db migration error: ", err)
	}

	linkRepository := postgres.NewLinkRepository(db)
	counterRepository := redis.NewCounterRepository(counter)

	createLinkUseCase := usecaselink.NewCreateLinkUseCase(linkRepository)
	getLinkUseCase := usecaselink.NewGetLinkUseCase(linkRepository, counterRepository)
	getCounterUseCase := usecasecounter.NewGetCounterUseCase(counterRepository)

	controller := web.NewLinkController(createLinkUseCase, getLinkUseCase, getCounterUseCase)

	router := gin.Default()

	router.POST("/api/v1/links", controller.CreateLink)
	router.GET("/api/v1/links/:code", controller.Forward)
	router.GET("/api/v1/links/:code/counter", controller.GetCount)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Printf("Server is running at %s", server.Addr)
		if err = server.ListenAndServe(); err != nil {
			log.Printf("Server stopped: %v", err)
		}
	}()

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	<-shutdown.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}
