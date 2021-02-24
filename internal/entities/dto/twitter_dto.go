package dto

type twitterResponse struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// TweetListDTO :nodoc
type TweetListDTO struct {
	Data []twitterResponse `json:"data"`
}
