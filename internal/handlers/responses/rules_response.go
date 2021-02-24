package responses

import (
	"net/http"

	"github.com/yurichandra/gunners/internal/entities/models"
)

// RulesItem :nodoc:
type RulesItem struct {
	Value string `json:"value"`
}

// RulesList :nodoc:
type RulesList struct {
	Data []RulesItem `json:"data"`
}

// Render :nodoc
func (r *RulesList) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

// Render :nodoc:
func (r *RulesItem) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

// NewRulesListResponse :nodoc:
func NewRulesListResponse(data []models.TwitterRules) *RulesList {
	var rulesItem []RulesItem

	for _, item := range data {
		newRulesItem := RulesItem{Value: item.Value}
		rulesItem = append(rulesItem, newRulesItem)
	}

	rulesList := &RulesList{
		Data: rulesItem,
	}

	return rulesList
}
