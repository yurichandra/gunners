package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/yurichandra/gunners/internal/services"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebsocketHandler :nodoc
type WebsocketHandler struct {
	twitterService *services.TwitterService
}

// NewWebsocketHandler :nodoc
func NewWebsocketHandler(twitterService *services.TwitterService) *WebsocketHandler {
	return &WebsocketHandler{
		twitterService: twitterService,
	}
}

// GetRoutes :nodoc
func (handler *WebsocketHandler) GetRoutes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", handler.Serve)

	return router
}

// Serve :nodoc
func (handler *WebsocketHandler) Serve(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		return
	}

	websocketService := services.NewWebsocketService(conn, *handler.twitterService)
	go websocketService.Serve(request.Context())
}
