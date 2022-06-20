package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3615121649203-3600651918983-sfUPfvbsryiEwKygDzj3ZC2K")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03JSUEERQ8-3615140230931-30ad9588ed666ae2e5c414889f265a464fdfe17b239736c8823d738fa0e3937b")

	// var bot = slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN" ))
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Example: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
				year := request.Param("year")
		 		yob, err := strconv.Atoi(year)
				 if err != nil{
					 fmt.Println("Error")
				 }
				 age := 2021-yob
				r := fmt.Sprintf("age is %d", age)
				response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}


// ===== notes =====

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"

// 	"github.com/shomali11/slacker"
// )

// // passing event to slackbot to print it out
// func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
// 	for event := range analyticsChannel{
// 		// prints to the console
// 		fmt.Println("Command Events")
// 		fmt.Println(event.Timestamp)
// 		fmt.Println(event.Command)
// 		fmt.Println(event.Parameters)
// 		fmt.Println(event.Event)
// 		fmt.Println()
// 	}
// }

// func main(){
// 	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3615121649203-3600651918983-sfUPfvbsryiEwKygDzj3ZC2K") // socket token
// 	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03JSUEERQ8-3615140230931-30ad9588ed666ae2e5c414889f265a464fdfe17b239736c8823d738fa0e3937b") // bot token

// 	// var bot = slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN" ))
// 	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

// 	go printCommandEvents(bot.CommandEvents()) // passes stuff to the slack command

// 	bot.Command("my yob is <year>", &slacker.CommandDefinition{
// 		Description: "yob calculator",
// 		Example: "my yob is 2020",
// 		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){  // slacker is coming from the slacker package
// 				year := request.Param("year")
// 		 		yob, err := strconv.Atoi(year) // converts the string year entered by the user into a number for calculation
// 				 if err != nil{
// 					 fmt.Println("Error")
// 				 }
// 				 age := 2022-yob
// 				r := fmt.Sprintf("Your age is %d", age)
// 				response.Reply(r)
// 		},
// 	})

// 	// REQUIRED for stop the bot
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	err := bot.Listen(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }