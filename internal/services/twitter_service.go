package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/yurichandra/gunners/internal/models"
	"github.com/yurichandra/gunners/internal/repositories"
)

// TwitterService :nodoc:
type TwitterService struct {
	http            *http.Client
	matchRepository repositories.MatchRepositoryContract
}

type twitterResponse struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type twitterListResponse struct {
	Data []twitterResponse `json:"data"`
}

// NewTwitterService :nodoc:
func NewTwitterService(client *http.Client, matchRepository repositories.MatchRepositoryContract) *TwitterService {
	return &TwitterService{
		http:            client,
		matchRepository: matchRepository,
	}
}

// GetRules :nodoc:
func (service *TwitterService) GetRules() ([]models.TwitterRules, error) {
	url := fmt.Sprintf("%s/2/tweets/search/stream/rules", os.Getenv("TWITTER_API_BASE"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []models.TwitterRules{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TWITTER_BEARER_TOKEN")))

	response, err := service.http.Do(req)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []models.TwitterRules{}, err
	}

	var rules twitterListResponse

	err = json.Unmarshal(body, &rules)
	if err != nil {
		return []models.TwitterRules{}, err
	}

	var twitterRulesList []models.TwitterRules
	for _, item := range rules.Data {
		twitterRules := models.TwitterRules{
			Value: item.Value,
		}

		twitterRulesList = append(twitterRulesList, twitterRules)
	}

	return twitterRulesList, nil
}

// SetRules :nodoc:
func (service *TwitterService) SetRules(rules []models.TwitterRules) (bool, error) {
	url := fmt.Sprintf("%s/2/tweets/search/stream/rules", os.Getenv("TWITTER_API_BASE"))

	body := createRequestBody(rules)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TWITTER_BEARER_TOKEN")))
	req.Header.Add("Content-Type", "application/json")

	response, err := service.http.Do(req)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	fmt.Println(response)

	return true, nil
}

// DetailStreamDataResponse :nodoc
type DetailStreamDataResponse struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// StreamDataResponse :nodoc
type StreamDataResponse struct {
	Data DetailStreamDataResponse `json:"data"`
}

// Stream :nodoc
func (service *TwitterService) Stream(ctx context.Context) {
	url := fmt.Sprintf("%s/2/tweets/search/stream", os.Getenv("TWITTER_API_BASE"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TWITTER_BEARER_TOKEN")))
	req.Header.Add("Content-Type", "application/json")

	newHTTP := &http.Client{}
	request := 0

	go func() {
		response, err := newHTTP.Do(req)
		request++
		if err != nil {
			log.Println(err.Error())
			return
		}

		for {
			switch response.StatusCode {
			case 200:
				fmt.Println(response)

				var stream StreamDataResponse

				err = json.NewDecoder(response.Body).Decode(&stream)
				if err != nil {
					fmt.Println(err.Error())
				}

				if stream.Data.Text != "" {
					service.handleReadData(ctx, stream.Data.Text)
				}
			case 429:
				fmt.Println(response)
				fmt.Println("You opened too many connections...")
			default:
				fmt.Println(response)
				fmt.Println(response.StatusCode)
			}
		}
	}()

	fmt.Println("[*] Listen tweet from Official FPL :D")
}

func (service *TwitterService) handleReadData(ctx context.Context, text string) {
	scorePattern := "\\d-\\d"
	tagPattern := "#[A-Z]+"

	newText := strings.Replace(text, "#FPL", "", 1)

	pattern, _ := regexp.Compile(scorePattern)
	match := pattern.MatchString(newText)
	fmt.Println(match)

	scoreString := pattern.FindString(newText)
	scores := strings.Split(scoreString, "-")

	newPattern := regexp.MustCompile(tagPattern)
	matchTag := newPattern.FindString(newText)

	ongoingMatch, _ := service.matchRepository.FindByTag(ctx, matchTag)
	pastScores := ongoingMatch.Score

	newHomeScore, _ := strconv.Atoi(scores[0])
	newAwayScore, _ := strconv.Atoi(scores[1])

	if newHomeScore > int(pastScores[0]) {
		fmt.Printf("%s is scored\n", ongoingMatch.HomeTeam)
	} else if newAwayScore > int(pastScores[1]) {
		fmt.Printf("%s is scored\n", ongoingMatch.AwayTeam)
	}

	ongoingMatch.Score = []uint{uint(newHomeScore), uint(newAwayScore)}
	service.matchRepository.Update(ctx, ongoingMatch)
}

func createRequestBody(rules []models.TwitterRules) []byte {
	var payloads []map[string]string
	for _, item := range rules {
		payload := map[string]string{
			"value": item.Value,
		}

		payloads = append(payloads, payload)
	}

	reqStruct := map[string]interface{}{
		"add": payloads,
	}

	reqString, _ := json.Marshal(reqStruct)
	return reqString
}
