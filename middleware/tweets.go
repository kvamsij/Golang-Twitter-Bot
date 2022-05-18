package middleware

import (
	"bet365/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func GetTweets(queryString string) models.Statuses {

	twitterUrl := GetUrl(queryString)
	credentials := getCredentials()
	method := "GET"
	var tweets models.Statuses
	client := &http.Client{}
	req, err := http.NewRequest(method, twitterUrl, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", credentials.Token))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	error := json.Unmarshal(body, &tweets)
	if error != nil {
		fmt.Println(error)
	}
	return tweets
}

func GetUrl(queryString string) string {

	isProd := getCredentials().IsProd
	twitterUrl := &url.URL{}
	queryParams := QueryParams(queryString)

	if !isProd {
		twitterUrl.Scheme = "http"
		twitterUrl.Host = "127.0.0.1:3000"
		twitterUrl.Path = "/Statuses"
	} else {
		twitterUrl.Scheme = "https"
		twitterUrl.Host = "api.twitter.com"
		twitterUrl.Path = "/1.1/search/tweets.json"
		twitterUrl.RawQuery = queryParams
	}
	return twitterUrl.String()
}

func QueryParams(queryString string) string {
	var queryParams string
	if strings.HasPrefix(queryString, "?") {
		queryParams = strings.Replace(queryString, "?", "", 1)
		queryString += "&tweet_mode=extended"
	} else {
		tempParams := &url.Values{}
		tempParams.Add("q", queryString)
		tempParams.Add("tweet_mode", "extended")
		tempParams.Add("count", "100")
		queryParams = tempParams.Encode()
	}
	return queryParams
}
