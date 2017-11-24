package domain

import "time"
import "fmt"

type Tweeter interface {
	PrintableTweet() string
}

type TextTweet struct {
	User string
	Text string
	Date *time.Time
	Id   int
	Read bool
}

func NewTextTweet(User string, Text string) *TextTweet {

	date := time.Now()

	t := TextTweet{
		User,
		Text,
		&date,
		0,
		false,
	}
	return &t
}

func (tweet *TextTweet) PrintableTweet() string {
	stringTweet := fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
	return stringTweet
}

func (tweet *TextTweet) String() string {
	return tweet.PrintableTweet()
}

type ImageTweet struct {
	TextTweet
	ImageUrl string
}

func NewImageTweet(User string, Text string, Image string) *ImageTweet {

	date := time.Now()

	t := ImageTweet{
		TextTweet{
			User,
			Text,
			&date,
			0,
			false,
		},
		Image,
	}
	return &t
}

func (tweet *ImageTweet) PrintableTweet() string {
	stringTweet := fmt.Sprintf("@%s: %s %s", tweet.User, tweet.Text, tweet.ImageUrl)
	return stringTweet
}

func (tweet *ImageTweet) String() string {
	return tweet.PrintableTweet()
}

type QuoteTweet struct {
	TextTweet
	Quote Tweeter
}

func NewQuoteTweet(User string, Text string, Quote Tweeter) *QuoteTweet {

	date := time.Now()

	t := QuoteTweet{
		TextTweet{
			User,
			Text,
			&date,
			0,
			false,
		},
		Quote,
	}
	return &t
}

func (tweet *QuoteTweet) PrintableTweet() string {
	stringTweet := fmt.Sprintf("@%s: %s \"%s\"", tweet.User, tweet.Text, tweet.Quote)
	return stringTweet
}

func (tweet *QuoteTweet) String() string {
	return tweet.PrintableTweet()
}
