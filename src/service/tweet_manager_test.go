package service_test

import (
	"testing"

	"github.com/Tweeter/src/domain"
	"github.com/Tweeter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	//Initialization
	var tweet *domain.Tweet

	user := "grupoEsfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	//Operation
	service.PublishTweet(tweet)

	//Validation
	publishedTweet := service.GetTweet()

	if publishedTweet.User != user && publishedTweet.Text != text {
		t.Error("Expected tweet is %s:%s \nbut is %s:%s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestCleanTweetDeletesTweet(t *testing.T) {

	//Initialization
	service.InitializeService()
	var tweet *domain.Tweet

	user := "grupoEsfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	//Operation
	service.PublishTweet(tweet)

	service.CleanTweets()

	if service.GetTweet() != nil {
		t.Error("Expected tweet is", tweet)
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	//Initialization
	service.InitializeService()
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	_, err = service.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected erro is user is required")
	}

}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	//Initialization
	service.InitializeService()
	var tweet *domain.Tweet

	user := "grupoesfera"
	var text string

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	_, err = service.PublishTweet(tweet)

	//Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	//Initialization
	service.InitializeService()
	var tweet *domain.Tweet

	user := "grupoesfera"
	text := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	_, err = service.PublishTweet(tweet)

	//Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text exceeds 140 characters")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet, secondTweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	// Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, 0, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, 1, user, secondText) {
		return
	}

}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCannotRetrieveTweetById(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweetById(15)

	if publishedTweet != nil {
		t.Errorf("Deberia devolver nil")
		return
	}
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)

	// Operation
	count := service.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user, text string) bool {

	if tweet.Id != id {
		t.Errorf("Expected id is %v but was %v", id, tweet.Id)
	}

	if tweet.User != user && tweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, tweet.User, tweet.Text)
		return false
	}

	if tweet.Date == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}
