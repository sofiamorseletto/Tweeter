package service

var tweet2 string

func PublishTweet(tweet string) {
	tweet2 = tweet
}

func GetTweet() string {
	return tweet2
}
