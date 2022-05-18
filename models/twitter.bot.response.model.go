package models

type BotResponse struct {
	SearchTerm   string `json:"searchTerm"`
	ResultCount  int    `json:"resultCount"`
	AvgWordCount int    `json:"avgWordCount"`
}
