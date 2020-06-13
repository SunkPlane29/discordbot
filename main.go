package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/SunkPlane29/discordbot/pkg/ytdl"
	dgo "github.com/bwmarrin/discordgo"
)

// Checking about pause and resume commands as well as
// creating queues
const (
	play    string = "!play"
	stop    string = "!stop"
	playUrl string = "!url"
	search  string = "!search"
	summon  string = "!summon"
	disband string = "!disband"
	help    string = "!help"
)

var commandsList []string = []string{play, stop, playUrl, search, summon, disband, help}

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

	fmt.Println("Bot connection with discord stabilished. Waiting for comamnds.")

	dg.AddHandler(messageCreate)
	var input string
	fmt.Scanln(&input)

}

func messageCreate(s *dgo.Session, m *dgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	fmt.Printf("New message from %s: %s\n", m.Author.Username, m.Content)

	message := strings.SplitN(m.Content, " ", 2)

	// Do something with those commands
	switch message[0] {
	case play:
		playCommand(s, m, message[1])
	case stop:
		// Stop all musics from playing
		reply(s, m, fmt.Sprintf("You typed %s", stop))
	case playUrl:
		// Instead of seachin a music download an direct url. Deal with error
		// handling.
		reply(s, m, fmt.Sprintf("You typed %s", playUrl))
	case search:
		// searches a list of five musics and display their titles. After that select
		// one music by indexing.
		reply(s, m, fmt.Sprintf("You typed %s", search))
	case summon:
		summonCommand(s, m)

	case disband:
		disbandCommand(s, m)
	case help:
		helpCommand(s, m)
	default:
		fmt.Println("No command")
	}
}

func reply(s *dgo.Session, m *dgo.MessageCreate, msg string) {
	s.ChannelMessageSend(m.ChannelID, msg)
}

func helpCommand(s *dgo.Session, m *dgo.MessageCreate) {

	reply(s, m, fmt.Sprintf("List of available commands: %s", commandsList))
}

func summonCommand(s *dgo.Session, m *dgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		// Make a reply function
		reply(s, m, "Could not find your channel")
		fmt.Println(err)
	}
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		reply(s, m, "Could not find your guild.")
		fmt.Println(err)
	}
	for _, vs := range g.VoiceStates {
		if vs.UserID == m.Author.ID {
			// Where false and true are muted and deaf states
			s.ChannelVoiceJoin(g.ID, vs.ChannelID, false, true)
		}
	}
}

func disbandCommand(s *dgo.Session, m *dgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		reply(s, m, "Could not find your channel.")
		fmt.Println(err)
	}
	vc, ok := s.VoiceConnections[c.GuildID]
	if !ok {
		reply(s, m, "Check if I am in a voice channel. If I am try the `!summon` command first.")
		fmt.Println(err)
		return
	}
	vc.Disconnect()

}

func playCommand(s *dgo.Session, m *dgo.MessageCreate, title string) {
	err := ytdl.DownloadDca(title, "./assets/audios/")
	if err != nil {
		reply(s, m, "Error while downloading the file. Check `!search` to see if your music exists.")
		fmt.Println(err)
		return
	}

	// Make a download audio and play audio func later

	reply(s, m, "Download successfull")

}
