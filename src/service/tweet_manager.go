package service

import (
	"fmt"
	"unicode/utf8"

	"github.com/Tweeter/src/domain"
)

var tweets []*domain.Tweet

func InitializeService() {

	tweets = make([]*domain.Tweet, 0)

}

func PublishTweet(tweet *domain.Tweet) error {
	if tweet.User == "" {
		return fmt.Errorf("user is required")
	}
	if tweet.Text == "" {
		return fmt.Errorf("text is required")
	}
	if utf8.RuneCountInString(tweet.Text) > 140 {
		return fmt.Errorf("text exceeds 140 characters")
	}
	tweets = append(tweets, tweet)
	return nil
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

func CleanTweets() {
	tweets = make([]*domain.Tweet, 0)
}
