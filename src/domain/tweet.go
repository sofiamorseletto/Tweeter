package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date *time.Time
}

func NewTweet(User string, Text string) *Tweet {

	date := time.Now()

	t := Tweet{
		User,
		Text,
		&date,
	}
	return &t
}
