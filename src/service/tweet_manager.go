package service

import "github.com/Tweeter/src/domain"

var tweet2 *domain.Tweet

func PublishTweet(tweet *domain.Tweet) {
	tweet2 = tweet
}

func GetTweet() *domain.Tweet {
	return tweet2
}

func CleanTweet() {
	tweet2 = nil
}
