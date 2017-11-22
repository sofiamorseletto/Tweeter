package service

import (
	"fmt"
	"unicode/utf8"

	"github.com/Tweeter/src/domain"
)

var tweet2 *domain.Tweet

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
	tweet2 = tweet
	return nil
}

func GetTweet() *domain.Tweet {
	return tweet2
}

func CleanTweet() {
	tweet2 = nil
}
