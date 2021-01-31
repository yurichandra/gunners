package requests

import (
	"errors"
	"net/http"
)

// RulesObjectRequest :nodoc:
type RulesObjectRequest struct {
	Value string `json:"value"`
}

// RulesRequest :nodoc:
type RulesRequest struct {
	Items []RulesObjectRequest
}

// Bind :nodoc:
func (r *RulesRequest) Bind(request *http.Request) error {
	if len(r.Items) == 0 {
		return errors.New("Value params should at least one")
	}

	if r.Items[0].Value == "" {
		return errors.New("Value can't be empty")
	}

	return nil
}
