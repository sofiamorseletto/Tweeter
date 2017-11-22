package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date *time.Time
	Id   int
}

func NewTweet(User string, Text string) *Tweet {

	date := time.Now()

	t := Tweet{
		User,
		Text,
		&date,
		0,
	}
	return &t
}
