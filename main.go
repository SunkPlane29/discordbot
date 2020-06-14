package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/SunkPlane29/discordbot/pkg/bot"
	dgo "github.com/bwmarrin/discordgo"
)

// Checking about pause and resume commands as well as
// creating queues
const (
	Play    string = "!play"
	Stop    string = "!stop"
	PlayUrl string = "!url"
	Search  string = "!search"
	Summon  string = "!summon"
	Disband string = "!disband"
	Help    string = "!help"
)

var token string

func init() {
	flag.StringVar(&token, "t", "", "Bot token")
	flag.Parse()
	if token == "" {
		fmt.Println("No token received.")
		os.Exit(1)
	}
}

func main() {

	dg, err := dgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	err = dg.Open()
	if err != nil {
		panic(err)
	}
	defer dg.Close()

	fmt.Println("Bot connection with discord stabilished. Waiting for comamnds.")

	dg.AddHandler(messageCreate)
	var input string
	fmt.Scanln(&input)

}

func messageCreate(s *dgo.Session, m *dgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	message := strings.SplitN(m.Content, " ", 2)

	// Do something with those commands
	switch message[0] {
	case Play:
		// Here I'm always passing s and m, couldn't I make a handler for this?
		bot.PlayCommand(s, m, message[1])

	case Stop:
		bot.Reply(s, m, "Command not yet supported.")

	case PlayUrl:
		bot.Reply(s, m, "Command not yet supported.")

	case Search:
		bot.Reply(s, m, "Command not yet supported")

	case Summon:
		bot.SummonCommand(s, m)

	case Disband:
		bot.DisbandCommand(s, m)

	case Help:
		bot.HelpCommand(s, m)
	}
}
