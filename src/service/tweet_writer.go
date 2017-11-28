package service

import (
	"os"

	"github.com/Tweeter/src/domain"
)

type TweetWriter interface {
	WriteTweet(chan domain.Tweeter, chan bool)
}

type MemoryTweetWriter struct {
	Tweets []domain.Tweeter
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	writer := MemoryTweetWriter{
		make([]domain.Tweeter, 0),
	}
	return &writer
}

func (writer *MemoryTweetWriter) WriteTweet(tweets chan domain.Tweeter, quit chan bool) {
	for tweet := range tweets {
		writer.Tweets = append(writer.Tweets, tweet)
	}
	quit <- true
}

type FileTweetWriter struct {
	file *os.File
}

func NewFileTweetWriter() *FileTweetWriter {
	file, _ := os.OpenFile(
		"tweets.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0600,
	)
	writer := new(FileTweetWriter)
	writer.file = file

	return writer
}

func (writer *FileTweetWriter) WriteTweet(tweet domain.Tweeter) {
	if writer.file != nil {
		byteSlice := []byte(tweet.PrintableTweet() + "\n")
		writer.file.Write(byteSlice)
	}
}
