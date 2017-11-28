package service

import "github.com/Tweeter/src/domain"

type ChannelTweetWriter struct {
	TweetWriter
}

func NewChannelTweetWriter(writer TweetWriter) *ChannelTweetWriter {
	channelWriter := ChannelTweetWriter{
		writer,
	}
	return &channelWriter
}

func (writer *ChannelTweetWriter) WriteTweet(tweets chan domain.Tweeter, quit chan bool) {
	writer.TweetWriter.WriteTweet(tweets, quit)
}
