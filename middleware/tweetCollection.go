package middleware

import (
	"bet365/models"
	"bet365/utils"
)

type TweetsCollection struct {
	Tweets []string
}

func NewTweetsCollection() TweetsCollection {
	collection := TweetsCollection{
		Tweets: []string{},
	}
	return collection
}

func (collection *TweetsCollection) AddTweetsToCollection(tweet string) {

	collection.Tweets = append(collection.Tweets, tweet)

}

func (collection *TweetsCollection) SanatizeData() {
	index := 0
	for _, tweet := range collection.Tweets {
		collection.Tweets[index] = utils.RemoveUrlsAndHashTagsFromString(tweet)
		index++
	}
}

func (collection *TweetsCollection) Compute() models.BotResponse {
	response := models.BotResponse{
		SearchTerm:   "",
		ResultCount:  len(collection.Tweets),
		AvgWordCount: 0,
	}
	if len(collection.Tweets) == 0 {
		return response
	}
	totalWordCount := 0
	for _, tweetText := range collection.Tweets {
		totalWordCount += len(tweetText)
	}

	response.AvgWordCount = totalWordCount / len(collection.Tweets)
	return response
}
