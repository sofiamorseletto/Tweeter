package service

import (
	"fmt"
	"unicode/utf8"

	"github.com/Tweeter/src/domain"
)

var tweets []*domain.Tweet
var users map[string]*domain.User
var id int

func InitializeService() {

	tweets = make([]*domain.Tweet, 0)
	users = make(map[string]*domain.User)
	id = 0

}

func PublishTweet(tweet *domain.Tweet) (int, error) {
	if tweet.User == "" {
		return 0, fmt.Errorf("user is required")
	}
	if tweet.Text == "" {
		return 0, fmt.Errorf("text is required")
	}
	if utf8.RuneCountInString(tweet.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}
	tweet.Id = id
	tweets = append(tweets, tweet)
	u, ok := users[tweet.User]

	if ok {
		u.Tweets = append(u.Tweets, tweet)
	} else {
		user := domain.NewUser(tweet.User)
		user.Tweets = append(user.Tweets, tweet)
		users[user.Name] = user
	}

	id = id + 1

	return tweet.Id, nil
}

func GetTweet() *domain.Tweet {
	if len(tweets) != 0 {
		return tweets[len(tweets)-1]
	}
	return nil
}

func GetTweets() []*domain.Tweet {
	return tweets
}

func GetTweetById(id int) *domain.Tweet {
	for _, tweet := range tweets {
		if tweet.Id == id {
			return tweet
		}
	}
	return nil
}

func CleanTweets() {
	tweets = make([]*domain.Tweet, 0)
	users = make(map[string]*domain.User)
	id = 0
}

func CountTweetsByUser(user string) int {
	return len(users[user].Tweets)
}

func GetTweetsByUser(user string) []*domain.Tweet {
	return users[user].Tweets
}

func Follow(user, userToFollow string) error {
	u1, ok := users[user]
	if !ok {
		return fmt.Errorf("User does not exist")
	}
	u2, ok2 := users[userToFollow]
	if !ok2 {
		return fmt.Errorf("User to follow does not exist")
	}
	u1.Following = append(u1.Following, u2)
	return nil
}

func GetTimeLine(user string) []*domain.Tweet {
	followingTweets := make([]*domain.Tweet, 0)
	u, ok := users[user]

	if ok {

		for _, following := range u.Following {
			followingTweets = append(followingTweets, following.Tweets...)
		}
		followingTweets = append(followingTweets, u.Tweets...)
	}

	return followingTweets
}
