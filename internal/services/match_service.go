package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/yurichandra/gunners/internal/entities/models"
	"github.com/yurichandra/gunners/internal/repositories"
)

// MatchService :nodoc
type MatchService struct {
	matchRepo      repositories.MatchRepositoryContract
	twitterService TwitterServiceContract
}

var officialFPLaccount = "OfficialFPL"

// NewMatchService :nodoc
func NewMatchService(
	matchRepository repositories.MatchRepositoryContract,
	twitterService TwitterServiceContract,
) *MatchService {
	return &MatchService{
		matchRepo:      matchRepository,
		twitterService: twitterService,
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

	err = service.storeRules(data)
	if err != nil {
		return models.Match{}, err
	}

	return service.matchRepo.Store(ctx, data)
}

func (service *MatchService) storeRules(data models.Match) error {
	ruleValue := fmt.Sprintf(
		"from:%s %s Goal",
		officialFPLaccount,
		data.Tag,
	)

	rules := make([]models.TwitterRules, 0)
	rules = append(rules, models.TwitterRules{Value: ruleValue})

	_, err := service.twitterService.SetRules(rules)
	if err != nil {
		return err
	}

	return nil
}
