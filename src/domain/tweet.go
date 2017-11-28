package domain

import "time"
import "fmt"

type Tweeter interface {
	GetUser() string
	GetText() string
	GetId() int
	GetDate() *time.Time
	GetRead() bool
	SetId(id int)
	SetRead(read bool)
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

func (tweet *TextTweet) GetId() int {
	return tweet.Id
}

func (tweet *TextTweet) GetUser() string {
	return tweet.User
}

func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

func (tweet *TextTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *TextTweet) GetRead() bool {
	return tweet.Read
}

func (tweet *TextTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *TextTweet) SetRead(read bool) {
	tweet.Read = read
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

func (tweet *ImageTweet) GetId() int {
	return tweet.Id
}

func (tweet *ImageTweet) GetUser() string {
	return tweet.User
}

func (tweet *ImageTweet) GetText() string {
	return tweet.Text
}

func (tweet *ImageTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *ImageTweet) GetRead() bool {
	return tweet.Read
}

func (tweet *ImageTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *ImageTweet) SetRead(read bool) {
	tweet.Read = read
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

func (tweet *QuoteTweet) GetId() int {
	return tweet.Id
}

func (tweet *QuoteTweet) GetUser() string {
	return tweet.User
}

func (tweet *QuoteTweet) GetText() string {
	return tweet.Text
}

func (tweet *QuoteTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *QuoteTweet) GetRead() bool {
	return tweet.Read
}

func (tweet *QuoteTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *QuoteTweet) SetRead(read bool) {
	tweet.Read = read
}
