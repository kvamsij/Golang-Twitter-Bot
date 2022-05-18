package models

type TwitterAPIResponse struct {
	CreatedAt string `json:"created_at"`
	ID        int64  `json:"id"`
	IDStr     string `json:"id_str"`
	FullText  string `json:"full_text"`
	Truncated bool   `json:"truncated"`
}

type SearchResultsMetadata struct {
	CompletedIn float64 `json:"completed_in"`
	MaxID       int64   `json:"max_id"`
	MaxIDStr    string  `json:"max_id_str"`
	NextResults string  `json:"next_results"`
	Query       string  `json:"query"`
	RefreshURL  string  `json:"refresh_url"`
	Count       int64   `json:"count"`
	SinceID     int64   `json:"since_id"`
	SinceIDStr  string  `json:"since_id_str"`
}

type Statuses struct {
	Statuses       []TwitterAPIResponse
	SearchMetadata SearchResultsMetadata `json:"search_metadata"`
}
