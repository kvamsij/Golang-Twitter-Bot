# Golang Twitter Bot

A bot to fetch tweets from Twitter (Twitter API) with provided search term and compute avarage number of words per tweet.

## Tasks achieved

- created an http server that run on port 8080 and only route is "/twitterbot" url to query for data
<http://localhost:8080/twitterbot?q=bet365>

- using Twitter API credentials, able to fetch tweets from search tweets endpoint  
- Sanitized data
  - removed single characters
  - removed emoji's
  - removed numeric values
  - removed urls from the tweet text

## required improvements

- Data cleaning can be improved using english dictionary data set to filter meaning full words.

## Limitations

- At the moment bot cannot fetch all the tweets from the endpoint so the logic was restricted to fetch about 500 tweets.

- following logics were tried to fetch tweets

  - to fetch tweets by requesting once

   ![15 tweets request](images/Screenshot%202022-05-18%20025230.png)

  - to fetch all tweets for a given query string using time sleep. Took about 15min with

    ```time.Sleep(8 * time.second)```
  
  ![all-tweets](/images/Screenshot%202022-05-18%20025037.png)

## How to

- rename .env.example to .env and paste BearerToken and set isProd to true

![set env](/images/Screenshot%202022-05-18%20033723.png)

- to run the Bot from the terminal in the root folder try ``` gin -p 8080 run . ```

## Bot sample response

![BotResponse](/images/Screenshot%202022-05-18%20023028.png)
