package services

import "github.com/yurichandra/gunners/internal/models"

// TwitterServiceContract :nodoc
type TwitterServiceContract interface {
	SetRules(rules []models.TwitterRules) (bool, error)
}
