package domain

import "fmt"

type TweetPlugin interface {
	ExecutePlugin() string
}

type FacebookPlugin struct{}

func (fb *FacebookPlugin) ExecutePlugin() string {
	return "Your tweet was shared on Facebook"
}

type CountTweetPlugin struct {
	tweets int
}

func (counter *CountTweetPlugin) ExecutePlugin() string {
	counter.tweets = counter.tweets + 1
	printLine := fmt.Sprintf("You have tweeted: %v tweets", counter.tweets)
	return printLine
}

func NewFacebookPlugin() *FacebookPlugin {
	fb := FacebookPlugin{}
	return &fb
}

func NewCountPlugin(counter int) *CountTweetPlugin {
	countPlugin := CountTweetPlugin{counter}
	return &countPlugin

}
