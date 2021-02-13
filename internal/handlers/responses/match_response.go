package responses

import (
	"net/http"

	"github.com/yurichandra/gunners/internal/models"
)

// MatchList :nodoc:
type MatchList struct {
	Data []models.Match `json:"data"`
}

// Render :nodoc
func (r *MatchList) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

// NewMatchListResponse :nodoc:
func NewMatchListResponse(data []models.Match) *MatchList {
	matchList := &MatchList{
		Data: data,
	}

	return matchList
}
