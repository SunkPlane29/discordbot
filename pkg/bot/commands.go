package bot

import (
	"fmt"

	"github.com/SunkPlane29/discordbot/pkg/youdl"
	dgo "github.com/bwmarrin/discordgo"
)

// Reply command.
func Reply(s *dgo.Session, m *dgo.MessageCreate, msg string) {
	s.ChannelMessageSend(m.ChannelID, msg)
}

// Summons the bot to the caller's voice channel
// Later making so that summon command return a vc type
func SummonCommand(s *dgo.Session, m *dgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		// Make a Reply function
		Reply(s, m, "Could not find your channel")
		fmt.Println(err)
	}
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		Reply(s, m, "Could not find your guild.")
		fmt.Println(err)
	}
	for _, vs := range g.VoiceStates {
		if vs.UserID == m.Author.ID {
			// Where false and true are muted and deaf states
			s.ChannelVoiceJoin(g.ID, vs.ChannelID, false, true)
		}
	}
}

// Disconnects the bot from it's current voice channel.
func DisbandCommand(s *dgo.Session, m *dgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		Reply(s, m, "Could not find your channel.")
		fmt.Println(err)
	}
	vc, ok := s.VoiceConnections[c.GuildID]
	if !ok {
		Reply(s, m, "Check if I am in a voice channel. If I am try the `!summon` command first.")
		fmt.Println(err)
		return
	}
	vc.Disconnect()

}

// Downloads the music files, encode to dca and later send them.
func PlayCommand(s *dgo.Session, m *dgo.MessageCreate, title string) {
	filedir := "./assets/audios/"
	filename, err := youdl.DownloadDca(title, filedir)
	if err != nil {
		Reply(s, m, "Error while downloading the file. Check `!search` to see if your music exists.")
		fmt.Println(err)
		return
	}

	// Make a download audio and play audio func later

	Reply(s, m, "Download successfull. Loading music.")

	SummonCommand(s, m)

	ch := make(chan []byte, 1024)
	doneCh := make(chan error)

	// One goroutine for handling loading the song and the other to send
	// the frames to the voice connection.
	go func(filepath string, ch chan []byte, doneChc chan error) {
		err := loadSong(filepath, ch, doneCh)
		if err != nil {
			fmt.Println(err)
		}
	}(filedir+filename, ch, doneCh)

	go func(s *dgo.Session, m *dgo.MessageCreate, ch chan []byte, doneCh chan error) {
		err := playSong(s, m, ch, doneCh)
		if err != nil {
			fmt.Println(err)
		}
	}(s, m, ch, doneCh)
}
