package model

// Response is the API response data.
type Response struct {
	Status  int      `json:"status"`
	Invoice *Invoice `json:"invoice,omitempty"`
	Message string   `json:"message,omitempty"`
}
