package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	dgo "github.com/bwmarrin/discordgo"
)

const (
	play    string = "!play"
	stop    string = "!stop"
	playUrl string = "!url"
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
	_ = dg
	err = dg.Open()
	if err != nil {
		panic(err)
	}
	defer dg.Close()
	dg.AddHandler(messageCreate)
	var input string
	fmt.Scanln(&input)

}

func messageCreate(s *dgo.Session, m *dgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	fmt.Printf("New message from %s: %s\n", m.Author.Username, m.Content)

	message := strings.Split(m.Content, " ")
	channelID := m.ChannelID

	// Do something with those commands
	switch message[0] {
	case play:
		fmt.Println("Sending message")
		s.ChannelMessageSend(channelID, "You typed !play")
	case stop:
		fmt.Println("Sending message")
		s.ChannelMessageSend(channelID, "You typed !stop")
	case playUrl:
		fmt.Println("Sending message")
		s.ChannelMessageSend(channelID, "You typed !url")
	}
}
