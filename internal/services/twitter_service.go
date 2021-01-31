package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/yurichandra/gunners/internal/models"
)

// TwitterService :nodoc:
type TwitterService struct {
	http *http.Client
}

type twitterResponse struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type twitterListResponse struct {
	Data []twitterResponse `json:"data"`
}

// NewTwitterService :nodoc:
func NewTwitterService(client *http.Client) *TwitterService {
	return &TwitterService{
		http: client,
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
