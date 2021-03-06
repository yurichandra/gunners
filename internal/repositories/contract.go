package repositories

import (
	"context"

	"github.com/yurichandra/gunners/internal/entities/models"
)

// MatchRepositoryContract :nodoc
type MatchRepositoryContract interface {
	Get(ctx context.Context) ([]models.Match, error)
	FindByTag(ctx context.Context, tag string) (models.Match, error)
	Store(ctx context.Context, data models.Match) (models.Match, error)
	Update(ctx context.Context, data models.Match) (models.Match, error)
}
