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

// List of commmands
const (
	Play    string = "!play"
	Summon  string = "!summon"
	Disband string = "!disband"
)

var token string

// Initializing with the token
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

	// Blocking command
	var input string
	fmt.Scanln(&input)

}

// Handler for message events, every time some member types a message
// this command get's called.
func messageCreate(s *dgo.Session, m *dgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	message := strings.SplitN(m.Content, " ", 2)

	switch message[0] {
	case Play:
		// Always passing s & m, optimization?
		bot.PlayCommand(s, m, message[1])

	case Summon:
		bot.SummonCommand(s, m)

	case Disband:
		bot.DisbandCommand(s, m)
	}
}
