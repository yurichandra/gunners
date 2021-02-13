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
	"github.com/yurichandra/gunners/internal/database"
	"github.com/yurichandra/gunners/internal/handlers"
	"github.com/yurichandra/gunners/internal/repositories"
	"github.com/yurichandra/gunners/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

var matchRepository *repositories.MatchRepository

var twitterService *services.TwitterService
var matchService *services.MatchService

var rulesHandler *handlers.RulesHandler
var matchHandler *handlers.MatchHandler

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

	db := database.InitDB()

	matchRepository = repositories.NewMatchRepository(db)

	twitterService = services.NewTwitterService(httpClient)
	matchService = services.NewMatchService(matchRepository, twitterService)

	rulesHandler = handlers.NewRulesHandler(twitterService)
	matchHandler = handlers.NewMatchHandler(matchService)
}

func initHTTP() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
	}))

	router.Mount("/matches", matchHandler.GetRoutes())
	router.Mount("/rules", rulesHandler.GetRoutes())

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Application is running on port :%s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
