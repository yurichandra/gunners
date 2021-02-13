package services

import (
	"context"

	"github.com/yurichandra/gunners/internal/models"
	"github.com/yurichandra/gunners/internal/repositories"
)

// MatchService :nodoc
type MatchService struct {
	matchRepo repositories.MatchRepositoryContract
}

// NewMatchService :nodoc
func NewMatchService(matchRepository repositories.MatchRepositoryContract) *MatchService {
	return &MatchService{
		matchRepo: matchRepository,
	}
}

// Get :nodoc
func (service *MatchService) Get(ctx context.Context) ([]models.Match, error) {
	return service.matchRepo.Get(ctx)
}
