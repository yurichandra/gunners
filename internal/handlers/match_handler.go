package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yurichandra/gunners/internal/handlers/responses"
	"github.com/yurichandra/gunners/internal/services"
)

// MatchHandler :nodoc
type MatchHandler struct {
	matchService services.MatchServiceContract
}

// NewMatchHandler :nodoc
func NewMatchHandler(matchService services.MatchServiceContract) *MatchHandler {
	return &MatchHandler{
		matchService: matchService,
	}
}

// GetRoutes :nodoc:
func (handler *MatchHandler) GetRoutes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", handler.GetMatches)

	return router
}

// GetMatches :nodoc:
func (handler *MatchHandler) GetMatches(writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	matches, err := handler.matchService.Get(ctx)
	if err != nil {
		errResponse := newErrorResponse(http.StatusInternalServerError, err.Error())
		render.Render(writer, request, errResponse)
		return
	}

	if err = render.Render(writer, request, responses.NewMatchListResponse(matches)); err != nil {
		errResponse := newErrorResponse(http.StatusInternalServerError, err.Error())
		render.Render(writer, request, errResponse)
		return
	}
}
