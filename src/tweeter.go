package main

import (
	"strconv"

	"github.com/Tweeter/src/domain"
	"github.com/Tweeter/src/service"
	"github.com/abiosoft/ishell"
)

func main() {
	shell := ishell.New()
	shell.SetPrompt("Tweeter >>")
	shell.Print("Type 'help' to know commands\n")
	service.InitializeService()

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {
			var tweet *domain.Tweet
			var id int

			defer c.ShowPrompt(true)

			c.Print("Write your user: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			tweet = domain.NewTweet(user, text)

			id, err := service.PublishTweet(tweet)
			if err != nil {
				c.Print("Error publishing tweet:", err)
			} else {
				c.Printf("Tweet sent with id: %v\n", id)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()

			if tweet != nil {
				c.Println(tweet.Text)
			} else {
				c.Println("")
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := service.GetTweets()

			if tweets != nil {
				for _, tweet := range tweets {
					c.Println(tweet.Text)
				}
			} else {
				c.Println("")
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "cleanTweets",
		Help: "Deletes previous tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			service.CleanTweets()

			c.Print("Tweets deleted\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "searchTweetsByUser",
		Help: "Searches tweets by user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Println("Enter the user: ")

			user := c.ReadLine()

			for _, tweet := range service.GetTweetsByUser(user) {
				c.Println(tweet.Text)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countTweetsByUser",
		Help: "Counts tweets by user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Println("Enter the user: ")

			user := c.ReadLine()

			c.Println(service.CountTweetsByUser(user))

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetById",
		Help: "Returns tweet with the same id",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Println("Enter the id: ")

			in := c.ReadLine()
			id, err := strconv.Atoi(in)

			if err != nil {
				// handle error
				c.Println("Invalid id")
			} else {
				c.Println(service.GetTweetById(id).Text)
			}
			return
		},
	})

	shell.Run()

}
