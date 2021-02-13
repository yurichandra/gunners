package requests

import (
	"errors"
	"net/http"
)

// MatchRequest :nodoc
type MatchRequest struct {
	HomeTeam string `json:"homeTeam"`
	AwayTeam string `json:"awayTeam"`
	Date     string `json:"date"`
}

// Bind :nodoc
func (r *MatchRequest) Bind(request *http.Request) error {
	if r.HomeTeam == "" {
		return errors.New("Field `homeTeam` is required")
	}

	if r.AwayTeam == "" {
		return errors.New("Field `awayTeam` is required")
	}

	if r.Date == "" {
		return errors.New("Field `date` is required")
	}

	return nil
}
