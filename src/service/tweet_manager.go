package service

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/Tweeter/src/domain"
)

type TweetManager struct {
	tweets []*domain.Tweet
	users  map[string]*domain.User
	words  map[string]int
	id     int
}

func NewTweetManager() *TweetManager {
	tm := TweetManager{
		make([]*domain.Tweet, 0),
		make(map[string]*domain.User),
		make(map[string]int),
		0,
	}
	return &tm
}

func (tweetManager *TweetManager) PublishTweet(tweet *domain.Tweet) (int, error) {
	if tweet.User == "" {
		return 0, fmt.Errorf("user is required")
	}
	if tweet.Text == "" {
		return 0, fmt.Errorf("text is required")
	}
	if utf8.RuneCountInString(tweet.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}
	tweet.Id = tweetManager.id
	tweetManager.tweets = append(tweetManager.tweets, tweet)
	u, ok := tweetManager.users[tweet.User]

	if ok {
		u.Tweets = append(u.Tweets, tweet)
	} else {
		user := domain.NewUser(tweet.User)
		user.Tweets = append(user.Tweets, tweet)
		tweetManager.users[user.Name] = user
	}

	tweetWords := strings.Fields(tweet.Text)
	for _, word := range tweetWords {
		tweetManager.words[word] = tweetManager.words[word] + 1
	}

	tweetManager.id++

	return tweet.Id, nil
}

func (tweetManager *TweetManager) GetTweet() *domain.Tweet {
	if len(tweetManager.tweets) != 0 {
		return tweetManager.tweets[len(tweetManager.tweets)-1]
	}
	return nil
}

func (tweetManager *TweetManager) GetTweets() []*domain.Tweet {
	return tweetManager.tweets
}

func (tweetManager *TweetManager) GetTweetById(id int) *domain.Tweet {
	for _, tweet := range tweetManager.tweets {
		if tweet.Id == id {
			return tweet
		}
	}
	return nil
}

func (tweetManager *TweetManager) CleanTweets() {
	tweetManager.tweets = make([]*domain.Tweet, 0)
	tweetManager.users = make(map[string]*domain.User)
	tweetManager.id = 0
}

func (tweetManager *TweetManager) CountTweetsByUser(user string) int {
	return len(tweetManager.users[user].Tweets)
}

func (tweetManager *TweetManager) GetTweetsByUser(user string) []*domain.Tweet {
	return tweetManager.users[user].Tweets
}

func (tweetManager *TweetManager) Follow(user, userToFollow string) error {
	u1, ok := tweetManager.users[user]
	if !ok {
		return fmt.Errorf("User does not exist")
	}
	u2, ok2 := tweetManager.users[userToFollow]
	if !ok2 {
		return fmt.Errorf("User to follow does not exist")
	}
	u1.Following = append(u1.Following, u2)
	return nil
}

func (tweetManager *TweetManager) GetTimeLine(user string) []*domain.Tweet {
	followingTweets := make([]*domain.Tweet, 0)
	u, ok := tweetManager.users[user]

	if ok {

		for _, following := range u.Following {
			followingTweets = append(followingTweets, following.Tweets...)
		}
		followingTweets = append(followingTweets, u.Tweets...)
	}

	return followingTweets
}

func (tweetManager *TweetManager) GetTrendingTopic() (string, string) {

}
