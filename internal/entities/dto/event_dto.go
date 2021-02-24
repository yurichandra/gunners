package dto

// EventDTO :nodoc
type EventDTO struct {
	MatchTag string `json:"matchTag"`
	Team     string `json:"team"`
	Event    string `json:"event"`
	Scores   []uint `json:"scores"`
}
