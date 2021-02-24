package services

import (
	"context"

	"github.com/yurichandra/gunners/internal/models"
)

// TwitterServiceContract :nodoc
type TwitterServiceContract interface {
	SetRules(rules []models.TwitterRules) (bool, error)
	GetRules() ([]models.TwitterRules, error)
	Stream(ctx context.Context)
}

// MatchServiceContract :nodoc
type MatchServiceContract interface {
	Get(ctx context.Context) ([]models.Match, error)
	Store(ctx context.Context, data models.Match) (models.Match, error)
}

// WebsocketServiceContract :nodoc
type WebsocketServiceContract interface {
	Serve(ctx context.Context)
}
