package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"github.com/yurichandra/gunners/internal/handlers"
	"github.com/yurichandra/gunners/internal/services"
)

var twitterService *services.TwitterService

var rulesHandler *handlers.RulesHandler

var httpClient http.Client

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Fail to load.env file")
	}
}

func bootstrap() {
	httpClient := &http.Client{
		Timeout: time.Second * 30,
	}

	twitterService = services.NewTwitterService(httpClient)
	rulesHandler = handlers.NewRulesHandler(twitterService)
}

func initHTTP() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
	}))

	router.Mount("/rules", rulesHandler.GetRoutes())

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Application is running on port :%s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
