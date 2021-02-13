package services

import (
	"context"

	"github.com/yurichandra/gunners/internal/models"
)

// TwitterServiceContract :nodoc
type TwitterServiceContract interface {
	SetRules(rules []models.TwitterRules) (bool, error)
	GetRules() ([]models.TwitterRules, error)
}

// MatchServiceContract :nodoc
type MatchServiceContract interface {
	Get(ctx context.Context) ([]models.Match, error)
}
