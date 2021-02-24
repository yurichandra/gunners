package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yurichandra/gunners/internal/entities/models"
	"github.com/yurichandra/gunners/internal/handlers/requests"
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
	router.Post("/", handler.StoreMatch)

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

// StoreMatch :nodoc
func (handler *MatchHandler) StoreMatch(writer http.ResponseWriter, request *http.Request) {
	var req requests.MatchRequest

	if err := render.Bind(request, &req); err != nil {
		errResponse := newErrorResponse(http.StatusUnprocessableEntity, err.Error())
		render.Render(writer, request, errResponse)
		return
	}

	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		errResponse := newErrorResponse(http.StatusBadRequest, err.Error())
		render.Render(writer, request, errResponse)
		return
	}

	data := models.Match{
		HomeTeam: req.HomeTeam,
		AwayTeam: req.AwayTeam,
		Date:     parsedDate,
	}

	ctx := context.Background()
	match, err := handler.matchService.Store(ctx, data)
	if err != nil {
		errResponse := newErrorResponse(http.StatusInternalServerError, err.Error())
		render.Render(writer, request, errResponse)
		return
	}

	if err = render.Render(writer, request, responses.NewMatchItemResponse(match)); err != nil {
		errResponse := newErrorResponse(http.StatusInternalServerError, err.Error())
		render.Render(writer, request, errResponse)
		return
	}
}
