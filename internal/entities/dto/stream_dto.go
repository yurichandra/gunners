package dto

// StreamDetailDTO :nodoc
type StreamDetailDTO struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// StreamDTO :nodoc
type StreamDTO struct {
	Data StreamDetailDTO `json:"data"`
}
