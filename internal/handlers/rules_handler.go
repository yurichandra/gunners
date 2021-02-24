package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yurichandra/gunners/internal/entities/models"
	"github.com/yurichandra/gunners/internal/handlers/requests"
	"github.com/yurichandra/gunners/internal/handlers/responses"
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

	router.Get("/", handler.GetRules)
	router.Post("/", handler.StoreRules)

	return router
}

// GetRules :nodoc:
func (handler *RulesHandler) GetRules(writer http.ResponseWriter, request *http.Request) {
	rules, err := handler.twitterService.GetRules()
	if err != nil {
		errResponse := newErrorResponse(http.StatusInternalServerError, err.Error())
		render.Render(writer, request, errResponse)
		return
	}

	if err = render.Render(writer, request, responses.NewRulesListResponse(rules)); err != nil {
		errResponse := newErrorResponse(http.StatusInternalServerError, err.Error())
		render.Render(writer, request, errResponse)
		return
	}
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
