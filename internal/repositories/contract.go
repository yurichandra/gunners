package repositories

import (
	"context"

	"github.com/yurichandra/gunners/internal/models"
)

// MatchRepositoryContract :nodoc
type MatchRepositoryContract interface {
	Get(ctx context.Context) ([]models.Match, error)
}
