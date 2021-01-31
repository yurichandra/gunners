package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yurichandra/gunners/internal/handlers/requests"
	"github.com/yurichandra/gunners/internal/models"
	"github.com/yurichandra/gunners/internal/services"
)

// RulesHandler :nodoc
type RulesHandler struct {
	twitterService services.TwitterServiceContract
}

// NewRulesHandler :nodoc
func NewRulesHandler(twitterService services.TwitterServiceContract) *RulesHandler {
	return &RulesHandler{
		twitterService: twitterService,
	}
}

// GetRoutes :nodoc:
func (handler *RulesHandler) GetRoutes() chi.Router {
	router := chi.NewRouter()

	router.Post("/", handler.StoreRules)

	return router
}

// StoreRules :nodoc:
func (handler *RulesHandler) StoreRules(writer http.ResponseWriter, request *http.Request) {
	var rulesRequest requests.RulesRequest

	if err := render.Bind(request, &rulesRequest); err != nil {
		errResponse := newErrorResponse(http.StatusUnprocessableEntity, err.Error())
		render.Render(writer, request, errResponse)
		return
	}

	var twitterRules []models.TwitterRules
	for _, item := range rulesRequest.Items {
		rule := models.TwitterRules{
			Value: item.Value,
		}

		twitterRules = append(twitterRules, rule)
	}

	_, err := handler.twitterService.SetRules(twitterRules)
	if err != nil {
		errResponse := newErrorResponse(http.StatusInternalServerError, err.Error())
		render.Render(writer, request, errResponse)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}
