package service

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/Tweeter/src/domain"
)

type TrendingTopic struct {
	word    string
	counter int
}

type TweetManager struct {
	tweets  []domain.Tweeter
	users   map[string]*domain.User
	words   map[string]int
	id      int
	topics  []*TrendingTopic
	plugins []domain.TweetPlugin
	channel *ChannelTweetWriter
}

func NewTweetManager(channel *ChannelTweetWriter) *TweetManager {
	tm := TweetManager{
		make([]domain.Tweeter, 0),
		make(map[string]*domain.User),
		make(map[string]int),
		0,
		make([]*TrendingTopic, 2),
		make([]domain.TweetPlugin, 0),
		channel,
	}
	tm.topics[0] = &TrendingTopic{"", 0}
	tm.topics[1] = &TrendingTopic{"", 0}

	return &tm
}

func (tweetManager *TweetManager) PublishTweet(tweet domain.Tweeter, quit chan bool) (int, error, []string) {
	if tweet.GetUser() == "" {
		return 0, fmt.Errorf("user is required"), nil
	}
	if tweet.GetText() == "" {
		return 0, fmt.Errorf("text is required"), nil
	}
	if utf8.RuneCountInString(tweet.GetText()) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters"), nil
	}

	tweetsToWrite := make(chan domain.Tweeter)

	go tweetManager.channel.WriteTweet(tweetsToWrite, quit)

	tweetsToWrite <- tweet

	close(tweetsToWrite)

	tweet.SetId(tweetManager.id)
	tweetManager.tweets = append(tweetManager.tweets, tweet)
	u, ok := tweetManager.users[tweet.GetUser()]

	if ok {
		u.Tweets = append(u.Tweets, tweet)
	} else {
		user := domain.NewUser(tweet.GetUser())
		user.Tweets = append(user.Tweets, tweet)
		tweetManager.users[user.Name] = user
	}

	tweetWords := strings.Fields(tweet.GetText())
	for _, word := range tweetWords {
		cant := tweetManager.words[word]
		cant = cant + 1
		tweetManager.words[word] = cant
		if cant > tweetManager.topics[0].counter && word != tweetManager.topics[0].word {
			w := TrendingTopic{
				word,
				cant,
			}
			tweetManager.topics[1] = tweetManager.topics[0]
			tweetManager.topics[0] = &w
		} else if cant > tweetManager.topics[1].counter && word != tweetManager.topics[0].word {
			w := TrendingTopic{
				word,
				cant,
			}
			tweetManager.topics[1] = &w
		}
	}

	tweetManager.id++

	return tweet.GetId(), nil, tweetManager.GetPlugins()
}

func (tweetManager *TweetManager) GetTweet() domain.Tweeter {
	if len(tweetManager.tweets) != 0 {
		return tweetManager.tweets[len(tweetManager.tweets)-1]
	}
	return nil
}

func (tweetManager *TweetManager) GetTweets() []domain.Tweeter {
	return tweetManager.tweets
}

func (tweetManager *TweetManager) GetTweetById(id int) domain.Tweeter {
	return tweetManager.getMessageById(id, tweetManager.tweets)
}

func (tweetManager *TweetManager) CleanTweets() {
	tweetManager.tweets = make([]domain.Tweeter, 0)
	tweetManager.users = make(map[string]*domain.User)
	tweetManager.id = 0
}

func (tweetManager *TweetManager) CountTweetsByUser(user string) int {
	foundUser, ok := tweetManager.users[user]
	if ok {
		return len(foundUser.Tweets)
	}
	return 0
}

func (tweetManager *TweetManager) GetTweetsByUser(user string) []domain.Tweeter {
	return tweetManager.users[user].Tweets
}

func (tweetManager *TweetManager) Follow(user, userToFollow string) error {
	u1, ok := tweetManager.users[user]
	if !ok {
		return fmt.Errorf("GetUser() does not exist")
	}
	u2, ok2 := tweetManager.users[userToFollow]
	if !ok2 {
		return fmt.Errorf("GetUser() to follow does not exist")
	}
	u1.Following = append(u1.Following, u2)
	return nil
}

func (tweetManager *TweetManager) GetTimeLine(user string) []domain.Tweeter {
	followingTweets := make([]domain.Tweeter, 0)
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
	return tweetManager.topics[0].word, tweetManager.topics[1].word
}

func (tweetManager *TweetManager) SendDirectMessage(user, userToMsg string, message domain.Tweeter) error {
	_, ok1 := tweetManager.users[user]
	_, ok2 := tweetManager.users[userToMsg]

	if ok1 && ok2 {
		tweetManager.users[userToMsg].DirectMessages = append(tweetManager.users[userToMsg].DirectMessages, message)
		return nil
	}

	return fmt.Errorf("The user does not exist")

}

func (tweetManager *TweetManager) GetAllDirectMessages(user string) []domain.Tweeter {
	return tweetManager.users[user].DirectMessages
}

func (tweetManager *TweetManager) ReadDirectMessage(user string, id int) {
	message := tweetManager.getMessageById(id, tweetManager.users[user].DirectMessages)
	message.SetRead(true)
}

func (tweetManager *TweetManager) getMessageById(id int, tweets []domain.Tweeter) domain.Tweeter {
	for _, tweet := range tweets {
		if tweet.GetId() == id {
			return tweet
		}
	}
	return nil
}

func (tm *TweetManager) AddPlugin(pluginToAdd domain.TweetPlugin) {
	tm.plugins = append(tm.plugins, pluginToAdd)
}

func (tw *TweetManager) GetPlugins() []string {
	pluginsList := make([]string, 0)
	for _, plugin := range tw.plugins {
		pluginsList = append(pluginsList, plugin.ExecutePlugin())
	}
	return pluginsList
}
