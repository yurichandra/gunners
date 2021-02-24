package cmd

import (
	"context"
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
var websocketHandler *handlers.WebsocketHandler

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

	twitterService = services.NewTwitterService(httpClient, matchRepository)
	matchService = services.NewMatchService(matchRepository, twitterService)

	rulesHandler = handlers.NewRulesHandler(twitterService)
	matchHandler = handlers.NewMatchHandler(matchService)
	websocketHandler = handlers.NewWebsocketHandler(twitterService)

	ctx := context.Background()

	// goroutine to handle listen stream.
	go twitterService.Stream(ctx)
}

func initHTTP() {
	httpRouter := chi.NewRouter()
	wsRouter := chi.NewRouter()

	httpRouter.Use(middleware.Logger)
	httpRouter.Use(render.SetContentType(render.ContentTypeJSON))
	httpRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	httpRouter.Mount("/matches", matchHandler.GetRoutes())
	httpRouter.Mount("/rules", rulesHandler.GetRoutes())

	wsRouter.Use(middleware.Logger)
	wsRouter.Use(render.SetContentType(render.ContentTypeJSON))
	wsRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))
	wsRouter.Mount("/ws", websocketHandler.GetRoutes())

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("HTTP protocol is running on port :%s\n", port)
	fmt.Printf("Websocket protocol is running on port 9000\n")

	// goroutine for handle websocket connection.
	go http.ListenAndServe(":9000", wsRouter)
	http.ListenAndServe(fmt.Sprintf(":%s", port), httpRouter)
}
