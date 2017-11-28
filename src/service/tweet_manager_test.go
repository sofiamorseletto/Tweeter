package service_test

import (
	"testing"

	"github.com/Tweeter/src/domain"
	"github.com/Tweeter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweeter

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	quit := make(chan bool)

	// Operation
	id, _, _ := tweetManager.PublishTweet(tweet, quit)

	// Validation
	publishedTweet := tweetManager.GetTweet()

	isValidTweet(t, publishedTweet, id, user, text)

	<-quit

	if memoryTweetWriter.Tweets[0] != tweet {
		t.Errorf("A tweet in the writer was expected")
	}

}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweeter

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	quit := make(chan bool)

	// Operation
	var err error
	_, err, _ = tweetManager.PublishTweet(tweet, quit)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweeter

	user := "grupoesfera"
	var text string

	tweet = domain.NewTextTweet(user, text)

	quit := make(chan bool)

	// Operation
	var err error
	_, err, _ = tweetManager.PublishTweet(tweet, quit)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweeter

	user := "grupoesfera"
	text := `The Go project has grown considerably with over half a million users and community members
	   all over the world. To date all community oriented activities have been organized by the community
	   with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet = domain.NewTextTweet(user, text)

	quit := make(chan bool)

	// Operation
	var err error
	_, err, _ = tweetManager.PublishTweet(tweet, quit)

	// Validation
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
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet domain.Tweeter

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)

	quit := make(chan bool)

	// Operation
	firstId, _, _ := tweetManager.PublishTweet(tweet, quit)
	secondId, _, _ := tweetManager.PublishTweet(secondTweet, quit)

	// Validation
	publishedTweets := tweetManager.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweeter
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	quit := make(chan bool)

	// Operation
	id, _, _ = tweetManager.PublishTweet(tweet, quit)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet, thirdTweet domain.Tweeter

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	quit := make(chan bool)

	tweetManager.PublishTweet(tweet, quit)
	tweetManager.PublishTweet(secondTweet, quit)
	tweetManager.PublishTweet(thirdTweet, quit)

	// Operation
	count := tweetManager.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet, thirdTweet domain.Tweeter

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	quit := make(chan bool)

	firstId, _, _ := tweetManager.PublishTweet(tweet, quit)
	secondId, _, _ := tweetManager.PublishTweet(secondTweet, quit)
	tweetManager.PublishTweet(thirdTweet, quit)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(tweets))
		return
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}

func isValidTweet(t *testing.T, tweet domain.Tweeter, id int, user, text string) bool {

	if tweet.GetId() != id {
		t.Errorf("Expected id is %v but was %v", id, tweet.GetId())
	}

	if tweet.GetUser() != user && tweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, tweet.GetUser(), tweet.GetText())
		return false
	}

	if tweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}

func TestFollowUser(t *testing.T) {
	//Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	tweet1 := domain.NewTextTweet("nportas", "Primer tweet")
	tweet2 := domain.NewTextTweet("mercadolibre", "Segundo tweet")
	tweet3 := domain.NewTextTweet("grupoesfera", "Tercer tweet")

	quit := make(chan bool)

	tweetManager.PublishTweet(tweet1, quit)
	tweetManager.PublishTweet(tweet2, quit)
	tweetManager.PublishTweet(tweet3, quit)

	//Operation
	tweetManager.Follow("grupoesfera", "nportas")
	tweetManager.Follow("grupoesfera", "mercadolibre")

	time_line := tweetManager.GetTimeLine("grupoesfera")

	if len(time_line) != 3 {
		t.Errorf("Expected lenght is 2 but was: %v", len(time_line))
		return
	}

	if !isValidTweet(t, time_line[0], 0, "nportas", "Primer tweet") {
		return
	}
	if !isValidTweet(t, time_line[1], 1, "mercadolibre", "Segundo tweet") {
		return
	}

}

func TestCanGetTrendingTopic(t *testing.T) {

	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet, thirdTweet domain.Tweeter

	user := "grupoesfera"
	anotherUser := "nick"
	text := "Hola va"
	secondText := "Hola como va"
	thirdText := "Hola va va"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, thirdText)

	quit := make(chan bool)

	tweetManager.PublishTweet(tweet, quit)
	tweetManager.PublishTweet(secondTweet, quit)
	tweetManager.PublishTweet(thirdTweet, quit)

	// Operation
	primera, segunda := tweetManager.GetTrendingTopic()

	// Validation
	if primera != "va" {
		t.Errorf("Expected word was va but is %s", primera)
		return
	}
	if segunda != "Hola" {
		t.Errorf("Expected word was Hola but is %s", segunda)
		return
	}

}

func TestSendDirectMessage(t *testing.T) {
	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet domain.Tweeter

	user := "grupoesfera"
	anotherUser := "nick"
	text := "Hola va"
	secondText := "Hola como va"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(anotherUser, secondText)

	quit := make(chan bool)

	tweetManager.PublishTweet(tweet, quit)
	tweetManager.PublishTweet(secondTweet, quit)

	//Operation
	tweetManager.SendDirectMessage(user, anotherUser, tweet)

	//Validation
	if len(tweetManager.GetAllDirectMessages(anotherUser)) != 1 {
		t.Errorf("Expected len of direct messages was 1 but is %v", len(tweetManager.GetAllDirectMessages(anotherUser)))
		return
	}

	if tweetManager.GetAllDirectMessages(anotherUser)[0].GetText() != "Hola va" {
		t.Errorf("Expected len of direct messages was Hola va but is %v", tweetManager.GetAllDirectMessages(anotherUser)[0].GetText())
		return
	}
}

func TestReadDirectMessage(t *testing.T) {
	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet domain.Tweeter

	user := "grupoesfera"
	anotherUser := "nick"
	text := "Hola va"
	secondText := "Hola como va"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(anotherUser, secondText)

	quit := make(chan bool)

	id, _, _ := tweetManager.PublishTweet(tweet, quit)
	tweetManager.PublishTweet(secondTweet, quit)

	//Operation
	tweetManager.SendDirectMessage(user, anotherUser, tweet)
	tweetManager.ReadDirectMessage(anotherUser, id)

	//Validation
	if tweetManager.GetAllDirectMessages(anotherUser)[0].GetRead() != true {
		t.Errorf("Expected status of messages was true but is %v", tweetManager.GetAllDirectMessages(anotherUser)[0].GetRead())
		return
	}

}

func TestPluginActivated(t *testing.T) {
	// Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetManager := service.NewTweetManager(tweetWriter)

	// var tweet domain.Tweeter
	var tweet2 domain.Tweeter

	user := "grupoesfera"
	// text := "Hola va"

	// tweet = domain.NewTextTweet(user, text)
	tweet2 = domain.NewTextTweet(user, "Hola")

	tweetManager.AddPlugin(domain.NewFacebookPlugin())
	tweetManager.AddPlugin(domain.NewCountPlugin(tweetManager.CountTweetsByUser(user)))

	quit := make(chan bool)

	//Operation
	_, _, plugins := tweetManager.PublishTweet(tweet2, quit)

	//Validation
	if len(plugins) == 0 {
		t.Errorf("Error: La longitud es: %v", len(plugins))
		return
	}

	if plugins[1] != "You have tweeted: 1 tweets" {
		t.Errorf("Expected status of messages was \"You have tweeted: 1 tweets\" but is %s", plugins[0])
		return
	}
	if plugins[0] != "Your tweet was shared on Facebook" {
		t.Errorf("Expected status of messages was \"Your tweet was shared on Facebook\" but is %s", plugins[1])
		return
	}
}

func TestCanWriteATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "Async tweet")
	tweet2 := domain.NewTextTweet("grupoesfera", "Async tweet2")

	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetsToWrite := make(chan domain.Tweeter)
	quit := make(chan bool)

	go tweetWriter.WriteTweet(tweetsToWrite, quit)

	// Operation
	tweetsToWrite <- tweet
	tweetsToWrite <- tweet2

	close(tweetsToWrite)
	<-quit

	// Validation
	if memoryTweetWriter.Tweets[0] != tweet {
		t.Errorf("A tweet in the writer was expected")
	}

	if memoryTweetWriter.Tweets[1] != tweet2 {
		t.Errorf("A tweet in the writer was expected")
	}
}
