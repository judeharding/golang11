https://www.youtube.com/watch?v=jFfo23yIWac @ 2.43

you need a slack account - https://api.slack.com/

you should put your tokens in a folder that does NOT get pushed up to github

you will need to work with your SOCKET MODE, EVENT SUBSCRIPTIONS and OAUTH & PERMISSIONS

go mod init github.com/judeharding/NAMEOFPROJECT
add a package = go git "github.com/shomali11/slacker"
then
go mod tidy // to clean up any outdated or duplicate packages

your project should have a go.mod and go.sum files which are like npm install package.json(initiatied by the go mod line in terminal)
create main.go file

when finished programming, run - go build - in terminal

HOW TO RUN

- in terminal go run main.go
- in slack type, @age-bot my yob is 1990
