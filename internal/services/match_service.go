package services

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

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

// Store :nodoc
func (service *MatchService) Store(ctx context.Context, data models.Match) (models.Match, error) {
	workdir, _ := os.Getwd()
	file, err := ioutil.ReadFile(workdir + "/internal/files/clubs.json")
	if err != nil {
		return models.Match{}, err
	}

	var club map[string]models.Club
	err = json.Unmarshal(file, &club)

	matchKey := "#" + club[data.HomeTeam].Abbreviation + club[data.AwayTeam].Abbreviation
	data.Tag = matchKey
	data.Score = []uint{0, 0}

	return service.matchRepo.Store(ctx, data)
}
