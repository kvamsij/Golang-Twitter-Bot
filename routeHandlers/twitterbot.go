package routeHandlers

import (
	"bet365/middleware"
	"bet365/models"
	"encoding/json"
	"net/http"
)

func TwitterBot(w http.ResponseWriter, r *http.Request) {

	var queryStringMap = r.URL.Query()
	tweetsCollection := middleware.NewTweetsCollection()
	var response models.BotResponse
	if !queryStringMap.Has("q") {
		json.NewEncoder(w).Encode(response)
		return
	}

	var tweets models.Statuses
	var queryString string = queryStringMap.Get("q")
	noOfSearchRequest := 5
	for i := 0; i < noOfSearchRequest; i++ {
		var twitterData = middleware.GetTweets(queryString)
		if len(twitterData.Statuses) == 0 {
			break
		}
		tweets.Statuses = append(tweets.Statuses, twitterData.Statuses...)
		queryString = twitterData.SearchMetadata.NextResults
	}

	for _, tweet := range tweets.Statuses {
		tweetsCollection.AddTweetsToCollection(tweet.FullText)
	}

	tweetsCollection.SanatizeData()
	response = tweetsCollection.Compute()
	response.SearchTerm = queryStringMap.Get("q")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
