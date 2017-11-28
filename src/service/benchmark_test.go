package service_test

import (
	"testing"

	"github.com/Tweeter/src/domain"
	"github.com/Tweeter/src/service"
)

func BenchmarkPublishTweetWithFileTweetWriter(b *testing.B) {

	// Initialization
	fileTweetWriter := service.NewFileTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(fileTweetWriter)
	tweetManager := service.NewTweetManager(tweetWriter)

	quit := make(chan bool)
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	for n := 0; n < b.N; n++ {
		tweetManager.PublishTweet(tweet, quit)
	}
}

func BenchmarkPublishTweetWithMemoryTweetWriter(b *testing.B) {

	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)
	tweetManager := service.NewTweetManager(tweetWriter)

	quit := make(chan bool)
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	for n := 0; n < b.N; n++ {
		tweetManager.PublishTweet(tweet, quit)
	}
}
