package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yurichandra/gunners/internal/models"
)

// TwitterService :nodoc:
type TwitterService struct {
	http *http.Client
}

// NewTwitterService :nodoc:
func NewTwitterService(client *http.Client) *TwitterService {
	return &TwitterService{
		http: client,
	}
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
