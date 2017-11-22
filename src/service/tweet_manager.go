package service

import (
	"fmt"

	"github.com/Tweeter/src/domain"
)

var tweet2 *domain.Tweet

func PublishTweet(tweet *domain.Tweet) error {
	if tweet.User == "" || tweet == nil {
		return fmt.Errorf("user is required")
	}
	tweet2 = tweet
	return nil
}

func GetTweet() *domain.Tweet {
	return tweet2
}

func CleanTweet() {
	tweet2 = nil
}
