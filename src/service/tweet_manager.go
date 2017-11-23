package service

import (
	"fmt"
	"unicode/utf8"

	"github.com/Tweeter/src/domain"
)

var tweets []*domain.Tweet
var tweets_map map[string][]*domain.Tweet
var id int

func InitializeService() {

	tweets = make([]*domain.Tweet, 0)
	tweets_map = make(map[string][]*domain.Tweet)
	id = 0

}

func PublishTweet(tweet *domain.Tweet) (int, error) {
	if tweet.User == "" {
		return 0, fmt.Errorf("user is required")
	}
	if tweet.Text == "" {
		return 0, fmt.Errorf("text is required")
	}
	if utf8.RuneCountInString(tweet.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}
	tweet.Id = id
	tweets = append(tweets, tweet)
	tweets_map[tweet.User] = append(tweets_map[tweet.User], tweet)
	id = id + 1

	return tweet.Id, nil
}

func GetTweet() *domain.Tweet {
	if len(tweets) != 0 {
		return tweets[len(tweets)-1]
	}
	return nil
}

func GetTweets() []*domain.Tweet {
	return tweets
}

func GetTweetById(id int) *domain.Tweet {
	for _, tweet := range tweets {
		if tweet.Id == id {
			return tweet
		}
	}
	return nil
}

func CleanTweets() {
	tweets = make([]*domain.Tweet, 0)
	tweets_map = make(map[string][]*domain.Tweet)
	id = 0
}

func CountTweetsByUser(user string) int {
	return len(tweets_map[user])
}

func GetTweetsByUser(user string) []*domain.Tweet {
	return tweets_map[user]
}
