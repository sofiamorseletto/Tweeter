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
	var tweet *domain.Tweet

	user := "grupoEsfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	//Operation
	service.PublishTweet(tweet)

	service.CleanTweet()

	if service.GetTweet() != nil {
		t.Error("Expected tweet is", tweet)
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	//Initialization
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	err = service.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected erro is user is required")
	}

}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	//Initialization
	var tweet *domain.Tweet

	user := "grupoesfera"
	var text string

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	err = service.PublishTweet(tweet)

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
	var tweet *domain.Tweet

	user := "grupoesfera"
	text := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	err = service.PublishTweet(tweet)

	//Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text exceeds 140 characters")
	}
}
